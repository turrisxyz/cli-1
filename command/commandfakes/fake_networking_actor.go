// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/cfnetworkingaction"
	v7 "code.cloudfoundry.org/cli/command"
)

type FakeNetworkingActor struct {
	AddNetworkPolicyStub        func(string, string, string, string, string, int, int) (cfnetworkingaction.Warnings, error)
	addNetworkPolicyMutex       sync.RWMutex
	addNetworkPolicyArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 int
		arg7 int
	}
	addNetworkPolicyReturns struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}
	addNetworkPolicyReturnsOnCall map[int]struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetworkingActor) AddNetworkPolicy(arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 int, arg7 int) (cfnetworkingaction.Warnings, error) {
	fake.addNetworkPolicyMutex.Lock()
	ret, specificReturn := fake.addNetworkPolicyReturnsOnCall[len(fake.addNetworkPolicyArgsForCall)]
	fake.addNetworkPolicyArgsForCall = append(fake.addNetworkPolicyArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
		arg5 string
		arg6 int
		arg7 int
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.recordInvocation("AddNetworkPolicy", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.addNetworkPolicyMutex.Unlock()
	if fake.AddNetworkPolicyStub != nil {
		return fake.AddNetworkPolicyStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.addNetworkPolicyReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNetworkingActor) AddNetworkPolicyCallCount() int {
	fake.addNetworkPolicyMutex.RLock()
	defer fake.addNetworkPolicyMutex.RUnlock()
	return len(fake.addNetworkPolicyArgsForCall)
}

func (fake *FakeNetworkingActor) AddNetworkPolicyCalls(stub func(string, string, string, string, string, int, int) (cfnetworkingaction.Warnings, error)) {
	fake.addNetworkPolicyMutex.Lock()
	defer fake.addNetworkPolicyMutex.Unlock()
	fake.AddNetworkPolicyStub = stub
}

func (fake *FakeNetworkingActor) AddNetworkPolicyArgsForCall(i int) (string, string, string, string, string, int, int) {
	fake.addNetworkPolicyMutex.RLock()
	defer fake.addNetworkPolicyMutex.RUnlock()
	argsForCall := fake.addNetworkPolicyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5, argsForCall.arg6, argsForCall.arg7
}

func (fake *FakeNetworkingActor) AddNetworkPolicyReturns(result1 cfnetworkingaction.Warnings, result2 error) {
	fake.addNetworkPolicyMutex.Lock()
	defer fake.addNetworkPolicyMutex.Unlock()
	fake.AddNetworkPolicyStub = nil
	fake.addNetworkPolicyReturns = struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworkingActor) AddNetworkPolicyReturnsOnCall(i int, result1 cfnetworkingaction.Warnings, result2 error) {
	fake.addNetworkPolicyMutex.Lock()
	defer fake.addNetworkPolicyMutex.Unlock()
	fake.AddNetworkPolicyStub = nil
	if fake.addNetworkPolicyReturnsOnCall == nil {
		fake.addNetworkPolicyReturnsOnCall = make(map[int]struct {
			result1 cfnetworkingaction.Warnings
			result2 error
		})
	}
	fake.addNetworkPolicyReturnsOnCall[i] = struct {
		result1 cfnetworkingaction.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworkingActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addNetworkPolicyMutex.RLock()
	defer fake.addNetworkPolicyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetworkingActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.NetworkingActor = new(FakeNetworkingActor)
