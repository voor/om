// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type StagedConfigService struct {
	GetDeployedProductCredentialStub        func(input api.GetDeployedProductCredentialInput) (api.GetDeployedProductCredentialOutput, error)
	getDeployedProductCredentialMutex       sync.RWMutex
	getDeployedProductCredentialArgsForCall []struct {
		input api.GetDeployedProductCredentialInput
	}
	getDeployedProductCredentialReturns struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}
	getDeployedProductCredentialReturnsOnCall map[int]struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}
	GetStagedProductByNameStub        func(product string) (api.StagedProductsFindOutput, error)
	getStagedProductByNameMutex       sync.RWMutex
	getStagedProductByNameArgsForCall []struct {
		product string
	}
	getStagedProductByNameReturns struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	getStagedProductByNameReturnsOnCall map[int]struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}
	GetStagedProductJobResourceConfigStub        func(productGUID, jobGUID string) (api.JobProperties, error)
	getStagedProductJobResourceConfigMutex       sync.RWMutex
	getStagedProductJobResourceConfigArgsForCall []struct {
		productGUID string
		jobGUID     string
	}
	getStagedProductJobResourceConfigReturns struct {
		result1 api.JobProperties
		result2 error
	}
	getStagedProductJobResourceConfigReturnsOnCall map[int]struct {
		result1 api.JobProperties
		result2 error
	}
	GetStagedProductNetworksAndAZsStub        func(product string) (map[string]interface{}, error)
	getStagedProductNetworksAndAZsMutex       sync.RWMutex
	getStagedProductNetworksAndAZsArgsForCall []struct {
		product string
	}
	getStagedProductNetworksAndAZsReturns struct {
		result1 map[string]interface{}
		result2 error
	}
	getStagedProductNetworksAndAZsReturnsOnCall map[int]struct {
		result1 map[string]interface{}
		result2 error
	}
	GetStagedProductPropertiesStub        func(product string) (map[string]api.ResponseProperty, error)
	getStagedProductPropertiesMutex       sync.RWMutex
	getStagedProductPropertiesArgsForCall []struct {
		product string
	}
	getStagedProductPropertiesReturns struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	getStagedProductPropertiesReturnsOnCall map[int]struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}
	ListDeployedProductsStub        func() ([]api.DeployedProductOutput, error)
	listDeployedProductsMutex       sync.RWMutex
	listDeployedProductsArgsForCall []struct{}
	listDeployedProductsReturns     struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	listDeployedProductsReturnsOnCall map[int]struct {
		result1 []api.DeployedProductOutput
		result2 error
	}
	ListStagedProductJobsStub        func(productGUID string) (map[string]string, error)
	listStagedProductJobsMutex       sync.RWMutex
	listStagedProductJobsArgsForCall []struct {
		productGUID string
	}
	listStagedProductJobsReturns struct {
		result1 map[string]string
		result2 error
	}
	listStagedProductJobsReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	ListStagedProductErrandsStub        func(productID string) (api.ErrandsListOutput, error)
	listStagedProductErrandsMutex       sync.RWMutex
	listStagedProductErrandsArgsForCall []struct {
		productID string
	}
	listStagedProductErrandsReturns struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	listStagedProductErrandsReturnsOnCall map[int]struct {
		result1 api.ErrandsListOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StagedConfigService) GetDeployedProductCredential(input api.GetDeployedProductCredentialInput) (api.GetDeployedProductCredentialOutput, error) {
	fake.getDeployedProductCredentialMutex.Lock()
	ret, specificReturn := fake.getDeployedProductCredentialReturnsOnCall[len(fake.getDeployedProductCredentialArgsForCall)]
	fake.getDeployedProductCredentialArgsForCall = append(fake.getDeployedProductCredentialArgsForCall, struct {
		input api.GetDeployedProductCredentialInput
	}{input})
	fake.recordInvocation("GetDeployedProductCredential", []interface{}{input})
	fake.getDeployedProductCredentialMutex.Unlock()
	if fake.GetDeployedProductCredentialStub != nil {
		return fake.GetDeployedProductCredentialStub(input)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDeployedProductCredentialReturns.result1, fake.getDeployedProductCredentialReturns.result2
}

func (fake *StagedConfigService) GetDeployedProductCredentialCallCount() int {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return len(fake.getDeployedProductCredentialArgsForCall)
}

func (fake *StagedConfigService) GetDeployedProductCredentialArgsForCall(i int) api.GetDeployedProductCredentialInput {
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	return fake.getDeployedProductCredentialArgsForCall[i].input
}

func (fake *StagedConfigService) GetDeployedProductCredentialReturns(result1 api.GetDeployedProductCredentialOutput, result2 error) {
	fake.GetDeployedProductCredentialStub = nil
	fake.getDeployedProductCredentialReturns = struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetDeployedProductCredentialReturnsOnCall(i int, result1 api.GetDeployedProductCredentialOutput, result2 error) {
	fake.GetDeployedProductCredentialStub = nil
	if fake.getDeployedProductCredentialReturnsOnCall == nil {
		fake.getDeployedProductCredentialReturnsOnCall = make(map[int]struct {
			result1 api.GetDeployedProductCredentialOutput
			result2 error
		})
	}
	fake.getDeployedProductCredentialReturnsOnCall[i] = struct {
		result1 api.GetDeployedProductCredentialOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductByName(product string) (api.StagedProductsFindOutput, error) {
	fake.getStagedProductByNameMutex.Lock()
	ret, specificReturn := fake.getStagedProductByNameReturnsOnCall[len(fake.getStagedProductByNameArgsForCall)]
	fake.getStagedProductByNameArgsForCall = append(fake.getStagedProductByNameArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetStagedProductByName", []interface{}{product})
	fake.getStagedProductByNameMutex.Unlock()
	if fake.GetStagedProductByNameStub != nil {
		return fake.GetStagedProductByNameStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductByNameReturns.result1, fake.getStagedProductByNameReturns.result2
}

func (fake *StagedConfigService) GetStagedProductByNameCallCount() int {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return len(fake.getStagedProductByNameArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductByNameArgsForCall(i int) string {
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	return fake.getStagedProductByNameArgsForCall[i].product
}

func (fake *StagedConfigService) GetStagedProductByNameReturns(result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	fake.getStagedProductByNameReturns = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductByNameReturnsOnCall(i int, result1 api.StagedProductsFindOutput, result2 error) {
	fake.GetStagedProductByNameStub = nil
	if fake.getStagedProductByNameReturnsOnCall == nil {
		fake.getStagedProductByNameReturnsOnCall = make(map[int]struct {
			result1 api.StagedProductsFindOutput
			result2 error
		})
	}
	fake.getStagedProductByNameReturnsOnCall[i] = struct {
		result1 api.StagedProductsFindOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfig(productGUID string, jobGUID string) (api.JobProperties, error) {
	fake.getStagedProductJobResourceConfigMutex.Lock()
	ret, specificReturn := fake.getStagedProductJobResourceConfigReturnsOnCall[len(fake.getStagedProductJobResourceConfigArgsForCall)]
	fake.getStagedProductJobResourceConfigArgsForCall = append(fake.getStagedProductJobResourceConfigArgsForCall, struct {
		productGUID string
		jobGUID     string
	}{productGUID, jobGUID})
	fake.recordInvocation("GetStagedProductJobResourceConfig", []interface{}{productGUID, jobGUID})
	fake.getStagedProductJobResourceConfigMutex.Unlock()
	if fake.GetStagedProductJobResourceConfigStub != nil {
		return fake.GetStagedProductJobResourceConfigStub(productGUID, jobGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductJobResourceConfigReturns.result1, fake.getStagedProductJobResourceConfigReturns.result2
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigCallCount() int {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return len(fake.getStagedProductJobResourceConfigArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigArgsForCall(i int) (string, string) {
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	return fake.getStagedProductJobResourceConfigArgsForCall[i].productGUID, fake.getStagedProductJobResourceConfigArgsForCall[i].jobGUID
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigReturns(result1 api.JobProperties, result2 error) {
	fake.GetStagedProductJobResourceConfigStub = nil
	fake.getStagedProductJobResourceConfigReturns = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductJobResourceConfigReturnsOnCall(i int, result1 api.JobProperties, result2 error) {
	fake.GetStagedProductJobResourceConfigStub = nil
	if fake.getStagedProductJobResourceConfigReturnsOnCall == nil {
		fake.getStagedProductJobResourceConfigReturnsOnCall = make(map[int]struct {
			result1 api.JobProperties
			result2 error
		})
	}
	fake.getStagedProductJobResourceConfigReturnsOnCall[i] = struct {
		result1 api.JobProperties
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZs(product string) (map[string]interface{}, error) {
	fake.getStagedProductNetworksAndAZsMutex.Lock()
	ret, specificReturn := fake.getStagedProductNetworksAndAZsReturnsOnCall[len(fake.getStagedProductNetworksAndAZsArgsForCall)]
	fake.getStagedProductNetworksAndAZsArgsForCall = append(fake.getStagedProductNetworksAndAZsArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetStagedProductNetworksAndAZs", []interface{}{product})
	fake.getStagedProductNetworksAndAZsMutex.Unlock()
	if fake.GetStagedProductNetworksAndAZsStub != nil {
		return fake.GetStagedProductNetworksAndAZsStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductNetworksAndAZsReturns.result1, fake.getStagedProductNetworksAndAZsReturns.result2
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsCallCount() int {
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	return len(fake.getStagedProductNetworksAndAZsArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsArgsForCall(i int) string {
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	return fake.getStagedProductNetworksAndAZsArgsForCall[i].product
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsReturns(result1 map[string]interface{}, result2 error) {
	fake.GetStagedProductNetworksAndAZsStub = nil
	fake.getStagedProductNetworksAndAZsReturns = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductNetworksAndAZsReturnsOnCall(i int, result1 map[string]interface{}, result2 error) {
	fake.GetStagedProductNetworksAndAZsStub = nil
	if fake.getStagedProductNetworksAndAZsReturnsOnCall == nil {
		fake.getStagedProductNetworksAndAZsReturnsOnCall = make(map[int]struct {
			result1 map[string]interface{}
			result2 error
		})
	}
	fake.getStagedProductNetworksAndAZsReturnsOnCall[i] = struct {
		result1 map[string]interface{}
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductProperties(product string) (map[string]api.ResponseProperty, error) {
	fake.getStagedProductPropertiesMutex.Lock()
	ret, specificReturn := fake.getStagedProductPropertiesReturnsOnCall[len(fake.getStagedProductPropertiesArgsForCall)]
	fake.getStagedProductPropertiesArgsForCall = append(fake.getStagedProductPropertiesArgsForCall, struct {
		product string
	}{product})
	fake.recordInvocation("GetStagedProductProperties", []interface{}{product})
	fake.getStagedProductPropertiesMutex.Unlock()
	if fake.GetStagedProductPropertiesStub != nil {
		return fake.GetStagedProductPropertiesStub(product)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStagedProductPropertiesReturns.result1, fake.getStagedProductPropertiesReturns.result2
}

func (fake *StagedConfigService) GetStagedProductPropertiesCallCount() int {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	return len(fake.getStagedProductPropertiesArgsForCall)
}

func (fake *StagedConfigService) GetStagedProductPropertiesArgsForCall(i int) string {
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	return fake.getStagedProductPropertiesArgsForCall[i].product
}

func (fake *StagedConfigService) GetStagedProductPropertiesReturns(result1 map[string]api.ResponseProperty, result2 error) {
	fake.GetStagedProductPropertiesStub = nil
	fake.getStagedProductPropertiesReturns = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) GetStagedProductPropertiesReturnsOnCall(i int, result1 map[string]api.ResponseProperty, result2 error) {
	fake.GetStagedProductPropertiesStub = nil
	if fake.getStagedProductPropertiesReturnsOnCall == nil {
		fake.getStagedProductPropertiesReturnsOnCall = make(map[int]struct {
			result1 map[string]api.ResponseProperty
			result2 error
		})
	}
	fake.getStagedProductPropertiesReturnsOnCall[i] = struct {
		result1 map[string]api.ResponseProperty
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListDeployedProducts() ([]api.DeployedProductOutput, error) {
	fake.listDeployedProductsMutex.Lock()
	ret, specificReturn := fake.listDeployedProductsReturnsOnCall[len(fake.listDeployedProductsArgsForCall)]
	fake.listDeployedProductsArgsForCall = append(fake.listDeployedProductsArgsForCall, struct{}{})
	fake.recordInvocation("ListDeployedProducts", []interface{}{})
	fake.listDeployedProductsMutex.Unlock()
	if fake.ListDeployedProductsStub != nil {
		return fake.ListDeployedProductsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDeployedProductsReturns.result1, fake.listDeployedProductsReturns.result2
}

func (fake *StagedConfigService) ListDeployedProductsCallCount() int {
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	return len(fake.listDeployedProductsArgsForCall)
}

func (fake *StagedConfigService) ListDeployedProductsReturns(result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	fake.listDeployedProductsReturns = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListDeployedProductsReturnsOnCall(i int, result1 []api.DeployedProductOutput, result2 error) {
	fake.ListDeployedProductsStub = nil
	if fake.listDeployedProductsReturnsOnCall == nil {
		fake.listDeployedProductsReturnsOnCall = make(map[int]struct {
			result1 []api.DeployedProductOutput
			result2 error
		})
	}
	fake.listDeployedProductsReturnsOnCall[i] = struct {
		result1 []api.DeployedProductOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductJobs(productGUID string) (map[string]string, error) {
	fake.listStagedProductJobsMutex.Lock()
	ret, specificReturn := fake.listStagedProductJobsReturnsOnCall[len(fake.listStagedProductJobsArgsForCall)]
	fake.listStagedProductJobsArgsForCall = append(fake.listStagedProductJobsArgsForCall, struct {
		productGUID string
	}{productGUID})
	fake.recordInvocation("ListStagedProductJobs", []interface{}{productGUID})
	fake.listStagedProductJobsMutex.Unlock()
	if fake.ListStagedProductJobsStub != nil {
		return fake.ListStagedProductJobsStub(productGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedProductJobsReturns.result1, fake.listStagedProductJobsReturns.result2
}

func (fake *StagedConfigService) ListStagedProductJobsCallCount() int {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return len(fake.listStagedProductJobsArgsForCall)
}

func (fake *StagedConfigService) ListStagedProductJobsArgsForCall(i int) string {
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	return fake.listStagedProductJobsArgsForCall[i].productGUID
}

func (fake *StagedConfigService) ListStagedProductJobsReturns(result1 map[string]string, result2 error) {
	fake.ListStagedProductJobsStub = nil
	fake.listStagedProductJobsReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductJobsReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.ListStagedProductJobsStub = nil
	if fake.listStagedProductJobsReturnsOnCall == nil {
		fake.listStagedProductJobsReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.listStagedProductJobsReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductErrands(productID string) (api.ErrandsListOutput, error) {
	fake.listStagedProductErrandsMutex.Lock()
	ret, specificReturn := fake.listStagedProductErrandsReturnsOnCall[len(fake.listStagedProductErrandsArgsForCall)]
	fake.listStagedProductErrandsArgsForCall = append(fake.listStagedProductErrandsArgsForCall, struct {
		productID string
	}{productID})
	fake.recordInvocation("ListStagedProductErrands", []interface{}{productID})
	fake.listStagedProductErrandsMutex.Unlock()
	if fake.ListStagedProductErrandsStub != nil {
		return fake.ListStagedProductErrandsStub(productID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listStagedProductErrandsReturns.result1, fake.listStagedProductErrandsReturns.result2
}

func (fake *StagedConfigService) ListStagedProductErrandsCallCount() int {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	return len(fake.listStagedProductErrandsArgsForCall)
}

func (fake *StagedConfigService) ListStagedProductErrandsArgsForCall(i int) string {
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	return fake.listStagedProductErrandsArgsForCall[i].productID
}

func (fake *StagedConfigService) ListStagedProductErrandsReturns(result1 api.ErrandsListOutput, result2 error) {
	fake.ListStagedProductErrandsStub = nil
	fake.listStagedProductErrandsReturns = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) ListStagedProductErrandsReturnsOnCall(i int, result1 api.ErrandsListOutput, result2 error) {
	fake.ListStagedProductErrandsStub = nil
	if fake.listStagedProductErrandsReturnsOnCall == nil {
		fake.listStagedProductErrandsReturnsOnCall = make(map[int]struct {
			result1 api.ErrandsListOutput
			result2 error
		})
	}
	fake.listStagedProductErrandsReturnsOnCall[i] = struct {
		result1 api.ErrandsListOutput
		result2 error
	}{result1, result2}
}

func (fake *StagedConfigService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDeployedProductCredentialMutex.RLock()
	defer fake.getDeployedProductCredentialMutex.RUnlock()
	fake.getStagedProductByNameMutex.RLock()
	defer fake.getStagedProductByNameMutex.RUnlock()
	fake.getStagedProductJobResourceConfigMutex.RLock()
	defer fake.getStagedProductJobResourceConfigMutex.RUnlock()
	fake.getStagedProductNetworksAndAZsMutex.RLock()
	defer fake.getStagedProductNetworksAndAZsMutex.RUnlock()
	fake.getStagedProductPropertiesMutex.RLock()
	defer fake.getStagedProductPropertiesMutex.RUnlock()
	fake.listDeployedProductsMutex.RLock()
	defer fake.listDeployedProductsMutex.RUnlock()
	fake.listStagedProductJobsMutex.RLock()
	defer fake.listStagedProductJobsMutex.RUnlock()
	fake.listStagedProductErrandsMutex.RLock()
	defer fake.listStagedProductErrandsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StagedConfigService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
