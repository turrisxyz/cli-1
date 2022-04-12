// Code generated by counterfeiter. DO NOT EDIT.
package sharedfakes

import (
	"sync"

	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv3/constant"
	"code.cloudfoundry.org/cli/command/shared"
	"code.cloudfoundry.org/cli/resources"
	"code.cloudfoundry.org/cli/util/configv3"
)

type FakeAppStager struct {
	StageAndStartStub        func(resources.Application, configv3.Space, configv3.Organization, string, constant.DeploymentStrategy, bool, constant.ApplicationAction) error
	stageAndStartMutex       sync.RWMutex
	stageAndStartArgsForCall []struct {
		arg1 resources.Application
		arg2 configv3.Space
		arg3 configv3.Organization
		arg4 string
		arg5 constant.DeploymentStrategy
		arg6 bool
		arg7 constant.ApplicationAction
	}
	stageAndStartReturns struct {
		result1 error
	}
	stageAndStartReturnsOnCall map[int]struct {
		result1 error
	}
	StageAppStub        func(resources.Application, string, configv3.Space) (resources.Droplet, error)
	stageAppMutex       sync.RWMutex
	stageAppArgsForCall []struct {
		arg1 resources.Application
		arg2 string
		arg3 configv3.Space
	}
	stageAppReturns struct {
		result1 resources.Droplet
		result2 error
	}
	stageAppReturnsOnCall map[int]struct {
		result1 resources.Droplet
		result2 error
	}
	StartAppStub        func(resources.Application, string, constant.DeploymentStrategy, bool, configv3.Space, configv3.Organization, constant.ApplicationAction) error
	startAppMutex       sync.RWMutex
	startAppArgsForCall []struct {
		arg1 resources.Application
		arg2 string
		arg3 constant.DeploymentStrategy
		arg4 bool
		arg5 configv3.Space
		arg6 configv3.Organization
		arg7 constant.ApplicationAction
	}
	startAppReturns struct {
		result1 error
	}
	startAppReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAppStager) StageAndStart(arg1 resources.Application, arg2 configv3.Space, arg3 configv3.Organization, arg4 string, arg5 constant.DeploymentStrategy, arg6 bool, arg7 constant.ApplicationAction) error {
	fake.stageAndStartMutex.Lock()
	ret, specificReturn := fake.stageAndStartReturnsOnCall[len(fake.stageAndStartArgsForCall)]
	fake.stageAndStartArgsForCall = append(fake.stageAndStartArgsForCall, struct {
		arg1 resources.Application
		arg2 configv3.Space
		arg3 configv3.Organization
		arg4 string
		arg5 constant.DeploymentStrategy
		arg6 bool
		arg7 constant.ApplicationAction
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("StageAndStart", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.stageAndStartMutex.Unlock()
	if fake.StageAndStartStub != nil {
		return fake.StageAndStartStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.stageAndStartReturns
	return fakeReturns.result1
}

func (fake *FakeAppStager) StageAndStartCallCount() int {
	fake.stageAndStartMutex.RLock()
	defer fake.stageAndStartMutex.RUnlock()
	return len(fake.stageAndStartArgsForCall)
}

func (fake *FakeAppStager) StageAndStartCalls(stub func(resources.Application, configv3.Space, configv3.Organization, string, constant.DeploymentStrategy, bool, constant.ApplicationAction) error) {
	fake.stageAndStartMutex.Lock()
	defer fake.stageAndStartMutex.Unlock()
	fake.StageAndStartStub = stub
}

func (fake *FakeAppStager) StageAndStartArgsForCall(i int) (resources.Application, configv3.Space, configv3.Organization, string, constant.DeploymentStrategy, bool, constant.ApplicationAction) {
	fake.stageAndStartMutex.RLock()
	defer fake.stageAndStartMutex.RUnlock()
	argsForCall := fake.stageAndStartArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeAppStager) StageAndStartReturns(result1 error) {
	fake.stageAndStartMutex.Lock()
	defer fake.stageAndStartMutex.Unlock()
	fake.StageAndStartStub = nil
	fake.stageAndStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppStager) StageAndStartReturnsOnCall(i int, result1 error) {
	fake.stageAndStartMutex.Lock()
	defer fake.stageAndStartMutex.Unlock()
	fake.StageAndStartStub = nil
	if fake.stageAndStartReturnsOnCall == nil {
		fake.stageAndStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stageAndStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppStager) StageApp(arg1 resources.Application, arg2 string, arg3 configv3.Space) (resources.Droplet, error) {
	fake.stageAppMutex.Lock()
	ret, specificReturn := fake.stageAppReturnsOnCall[len(fake.stageAppArgsForCall)]
	fake.stageAppArgsForCall = append(fake.stageAppArgsForCall, struct {
		arg1 resources.Application
		arg2 string
		arg3 configv3.Space
	}{arg1, arg2, arg3})
	fake.recordInvocation("StageApp", []interface{}{arg1, arg2, arg3})
	fake.stageAppMutex.Unlock()
	if fake.StageAppStub != nil {
		return fake.StageAppStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.stageAppReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAppStager) StageAppCallCount() int {
	fake.stageAppMutex.RLock()
	defer fake.stageAppMutex.RUnlock()
	return len(fake.stageAppArgsForCall)
}

func (fake *FakeAppStager) StageAppCalls(stub func(resources.Application, string, configv3.Space) (resources.Droplet, error)) {
	fake.stageAppMutex.Lock()
	defer fake.stageAppMutex.Unlock()
	fake.StageAppStub = stub
}

func (fake *FakeAppStager) StageAppArgsForCall(i int) (resources.Application, string, configv3.Space) {
	fake.stageAppMutex.RLock()
	defer fake.stageAppMutex.RUnlock()
	argsForCall := fake.stageAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeAppStager) StageAppReturns(result1 resources.Droplet, result2 error) {
	fake.stageAppMutex.Lock()
	defer fake.stageAppMutex.Unlock()
	fake.StageAppStub = nil
	fake.stageAppReturns = struct {
		result1 resources.Droplet
		result2 error
	}{result1, result2}
}

func (fake *FakeAppStager) StageAppReturnsOnCall(i int, result1 resources.Droplet, result2 error) {
	fake.stageAppMutex.Lock()
	defer fake.stageAppMutex.Unlock()
	fake.StageAppStub = nil
	if fake.stageAppReturnsOnCall == nil {
		fake.stageAppReturnsOnCall = make(map[int]struct {
			result1 resources.Droplet
			result2 error
		})
	}
	fake.stageAppReturnsOnCall[i] = struct {
		result1 resources.Droplet
		result2 error
	}{result1, result2}
}

func (fake *FakeAppStager) StartApp(arg1 resources.Application, arg2 string, arg3 constant.DeploymentStrategy, arg4 bool, arg5 configv3.Space, arg6 configv3.Organization, arg7 constant.ApplicationAction) error {
	fake.startAppMutex.Lock()
	ret, specificReturn := fake.startAppReturnsOnCall[len(fake.startAppArgsForCall)]
	fake.startAppArgsForCall = append(fake.startAppArgsForCall, struct {
		arg1 resources.Application
		arg2 string
		arg3 constant.DeploymentStrategy
		arg4 bool
		arg5 configv3.Space
		arg6 configv3.Organization
		arg7 constant.ApplicationAction
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("StartApp", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.startAppMutex.Unlock()
	if fake.StartAppStub != nil {
		return fake.StartAppStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.startAppReturns
	return fakeReturns.result1
}

func (fake *FakeAppStager) StartAppCallCount() int {
	fake.startAppMutex.RLock()
	defer fake.startAppMutex.RUnlock()
	return len(fake.startAppArgsForCall)
}

func (fake *FakeAppStager) StartAppCalls(stub func(resources.Application, string, constant.DeploymentStrategy, bool, configv3.Space, configv3.Organization, constant.ApplicationAction) error) {
	fake.startAppMutex.Lock()
	defer fake.startAppMutex.Unlock()
	fake.StartAppStub = stub
}

func (fake *FakeAppStager) StartAppArgsForCall(i int) (resources.Application, string, constant.DeploymentStrategy, bool, configv3.Space, configv3.Organization, constant.ApplicationAction) {
	fake.startAppMutex.RLock()
	defer fake.startAppMutex.RUnlock()
	argsForCall := fake.startAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeAppStager) StartAppReturns(result1 error) {
	fake.startAppMutex.Lock()
	defer fake.startAppMutex.Unlock()
	fake.StartAppStub = nil
	fake.startAppReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppStager) StartAppReturnsOnCall(i int, result1 error) {
	fake.startAppMutex.Lock()
	defer fake.startAppMutex.Unlock()
	fake.StartAppStub = nil
	if fake.startAppReturnsOnCall == nil {
		fake.startAppReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startAppReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAppStager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stageAndStartMutex.RLock()
	defer fake.stageAndStartMutex.RUnlock()
	fake.stageAppMutex.RLock()
	defer fake.stageAppMutex.RUnlock()
	fake.startAppMutex.RLock()
	defer fake.startAppMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAppStager) recordInvocation(key string, args []interface{}) {
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

var _ shared.AppStager = new(FakeAppStager)
