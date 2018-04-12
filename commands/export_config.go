package commands

import (
	"fmt"

	"github.com/pivotal-cf/jhanda"
	"github.com/pivotal-cf/om/api"
	yaml "gopkg.in/yaml.v2"
)

type ExportConfig struct {
	logger              logger
	exportConfigService exportConfigService
	Options             struct {
		Product string `long:"product-name"    short:"p" required:"true" description:"name of product"`
	}
}

//go:generate counterfeiter -o ./fakes/export_config_service.go --fake-name ExportConfigService . exportConfigService
type exportConfigService interface {
	Find(product string) (api.StagedProductsFindOutput, error)
	Jobs(productGUID string) (map[string]string, error)
	GetExistingJobConfig(productGUID, jobGUID string) (api.JobProperties, error)
	Properties(product string) (map[string]api.ResponseProperty, error)
	NetworksAndAZs(product string) (map[string]interface{}, error)
}

func NewExportConfig(exportConfigService exportConfigService, logger logger) ExportConfig {
	return ExportConfig{
		logger:              logger,
		exportConfigService: exportConfigService,
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

	findOutput, err := ec.exportConfigService.Find(ec.Options.Product)
	if err != nil {
		return err
	}
	productGUID := findOutput.Product.GUID

	properties, err := ec.exportConfigService.Properties(productGUID)
	if err != nil {
		return err
	}

	configurableProperties := map[string]interface{}{}

	for name, property := range properties {
		if property.Configurable {
			configurableProperties[name] = map[string]interface{}{"value": property.Value}
		}
	}

	networks, err := ec.exportConfigService.NetworksAndAZs(productGUID)
	if err != nil {
		return err
	}

	jobs, err := ec.exportConfigService.Jobs(productGUID)
	if err != nil {
		return err
	}

	resourceConfig := map[string]api.JobProperties{}

	for name, jobGUID := range jobs {
		jobProperties, err := ec.exportConfigService.GetExistingJobConfig(productGUID, jobGUID)
		if err != nil {
			return err
		}

		resourceConfig[name] = jobProperties
	}

	config := struct {
		Properties               map[string]interface{}       `yaml:"product-properties"`
		NetworkProperties        map[string]interface{}       `yaml:"network-properties"`
		ResourceConfigProperties map[string]api.JobProperties `yaml:"resource-config"`
	}{
		Properties:               configurableProperties,
		NetworkProperties:        networks,
		ResourceConfigProperties: resourceConfig,
	}

	output, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %s", err) // un-tested
	}
	ec.logger.Println(string(output))

	return nil
}
