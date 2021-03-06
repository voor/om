// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type DeployedProductsService struct {
	GetDiagnosticReportStub        func() (api.DiagnosticReport, error)
	getDiagnosticReportMutex       sync.RWMutex
	getDiagnosticReportArgsForCall []struct{}
	getDiagnosticReportReturns     struct {
		result1 api.DiagnosticReport
		result2 error
	}
	getDiagnosticReportReturnsOnCall map[int]struct {
		result1 api.DiagnosticReport
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeployedProductsService) GetDiagnosticReport() (api.DiagnosticReport, error) {
	fake.getDiagnosticReportMutex.Lock()
	ret, specificReturn := fake.getDiagnosticReportReturnsOnCall[len(fake.getDiagnosticReportArgsForCall)]
	fake.getDiagnosticReportArgsForCall = append(fake.getDiagnosticReportArgsForCall, struct{}{})
	fake.recordInvocation("GetDiagnosticReport", []interface{}{})
	fake.getDiagnosticReportMutex.Unlock()
	if fake.GetDiagnosticReportStub != nil {
		return fake.GetDiagnosticReportStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDiagnosticReportReturns.result1, fake.getDiagnosticReportReturns.result2
}

func (fake *DeployedProductsService) GetDiagnosticReportCallCount() int {
	fake.getDiagnosticReportMutex.RLock()
	defer fake.getDiagnosticReportMutex.RUnlock()
	return len(fake.getDiagnosticReportArgsForCall)
}

func (fake *DeployedProductsService) GetDiagnosticReportReturns(result1 api.DiagnosticReport, result2 error) {
	fake.GetDiagnosticReportStub = nil
	fake.getDiagnosticReportReturns = struct {
		result1 api.DiagnosticReport
		result2 error
	}{result1, result2}
}

func (fake *DeployedProductsService) GetDiagnosticReportReturnsOnCall(i int, result1 api.DiagnosticReport, result2 error) {
	fake.GetDiagnosticReportStub = nil
	if fake.getDiagnosticReportReturnsOnCall == nil {
		fake.getDiagnosticReportReturnsOnCall = make(map[int]struct {
			result1 api.DiagnosticReport
			result2 error
		})
	}
	fake.getDiagnosticReportReturnsOnCall[i] = struct {
		result1 api.DiagnosticReport
		result2 error
	}{result1, result2}
}

func (fake *DeployedProductsService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDiagnosticReportMutex.RLock()
	defer fake.getDiagnosticReportMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeployedProductsService) recordInvocation(key string, args []interface{}) {
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
