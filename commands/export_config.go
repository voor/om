package commands

import (
	"fmt"

	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/api"
	yaml "gopkg.in/yaml.v2"
)

type ExportConfig struct {
	logger                logger
	exportConfigService   exportConfigService
	exportJobsService     jobsService
	stagedProductsService stagedProductsService
	Options               struct {
		Product string `long:"product-name"    short:"p" required:"true" description:"name of product"`
	}
}

//go:generate counterfeiter -o ./fakes/export_config_service.go --fake-name ExportConfigService . exportConfigService
type exportConfigService interface {
	ExportConfig(product string) (api.ExportConfigOutput, error)
}

//go:generate counterfeiter -o ./fakes/export_jobs_config.go --fake-name ExportJobsConfig . exportJobsConfig
type exportJobsConfig interface {
	Jobs(productGUID string) (map[string]string, error)
	GetExistingJobConfig(productGUID, jobGUID string) (api.JobProperties, error)
}

type omConfigOutput struct {
	Properties               map[string]api.OutputProperty `yaml:"product-properties"`
	NetworkProperties        map[string]interface{}        `yaml:"network-properties"`
	ResourceConfigProperties map[string]api.JobProperties  `yaml:"resource-config"`
}

func NewExportConfig(exportConfigService exportConfigService, exportJobsService jobsService, stagedProductsService stagedProductsService, logger logger) ExportConfig {
	return ExportConfig{
		logger:                logger,
		exportConfigService:   exportConfigService,
		exportJobsService:     exportJobsService,
		stagedProductsService: stagedProductsService,
	}
}

func (ec ExportConfig) Usage() jhanda.Usage {
	return jhanda.Usage{
		Description:      "This command generates a config from a staged product that can be passed in to om configure-product",
		ShortDescription: "generates a config from a staged product",
		Flags:            ec.Options,
	}
}

func (ec ExportConfig) Execute(args []string) error {
	if _, err := jhanda.Parse(&ec.Options, args); err != nil {
		return fmt.Errorf("could not parse export-config flags: %s", err)
	}

	productConfig, err := ec.exportConfigService.ExportConfig(ec.Options.Product)
	if err != nil {
		return fmt.Errorf("failed to export config: %s", err)
	}

	findOutput, err := ec.stagedProductsService.Find(ec.Options.Product)
	if err != nil {
		return fmt.Errorf("could not find staged product with name 'p-bosh': %s", err)
	}
	productGUID := findOutput.Product.GUID

	jobs, err := ec.exportJobsService.Jobs(productGUID)
	if err != nil {
		return fmt.Errorf("failed to fetch jobs: %s", err)
	}

	resourceConfig := map[string]api.JobProperties{}

	for name, jobGUID := range jobs {
		jobProperties, err := ec.exportJobsService.GetExistingJobConfig(productGUID, jobGUID)
		if err != nil {
			panic(err)
		}

		resourceConfig[name] = jobProperties
	}

	config := omConfigOutput{
		Properties:               productConfig.Properties,
		NetworkProperties:        productConfig.NetworkProperties,
		ResourceConfigProperties: resourceConfig,
	}

	output, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err) // un-tested
	}
	ec.logger.Println(string(output))

	return nil
}
