// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/command/v2"
)

type FakeUninstallPluginActor struct {
	UninstallPluginStub        func(name string) (v2action.Warnings, error)
	uninstallPluginMutex       sync.RWMutex
	uninstallPluginArgsForCall []struct {
		name string
	}
	uninstallPluginReturns struct {
		result1 v2action.Warnings
		result2 error
	}
	uninstallPluginReturnsOnCall map[int]struct {
		result1 v2action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUninstallPluginActor) UninstallPlugin(name string) (v2action.Warnings, error) {
	fake.uninstallPluginMutex.Lock()
	ret, specificReturn := fake.uninstallPluginReturnsOnCall[len(fake.uninstallPluginArgsForCall)]
	fake.uninstallPluginArgsForCall = append(fake.uninstallPluginArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("UninstallPlugin", []interface{}{name})
	fake.uninstallPluginMutex.Unlock()
	if fake.UninstallPluginStub != nil {
		return fake.UninstallPluginStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.uninstallPluginReturns.result1, fake.uninstallPluginReturns.result2
}

func (fake *FakeUninstallPluginActor) UninstallPluginCallCount() int {
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	return len(fake.uninstallPluginArgsForCall)
}

func (fake *FakeUninstallPluginActor) UninstallPluginArgsForCall(i int) string {
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	return fake.uninstallPluginArgsForCall[i].name
}

func (fake *FakeUninstallPluginActor) UninstallPluginReturns(result1 v2action.Warnings, result2 error) {
	fake.UninstallPluginStub = nil
	fake.uninstallPluginReturns = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUninstallPluginActor) UninstallPluginReturnsOnCall(i int, result1 v2action.Warnings, result2 error) {
	fake.UninstallPluginStub = nil
	if fake.uninstallPluginReturnsOnCall == nil {
		fake.uninstallPluginReturnsOnCall = make(map[int]struct {
			result1 v2action.Warnings
			result2 error
		})
	}
	fake.uninstallPluginReturnsOnCall[i] = struct {
		result1 v2action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeUninstallPluginActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uninstallPluginMutex.RLock()
	defer fake.uninstallPluginMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUninstallPluginActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.UninstallPluginActor = new(FakeUninstallPluginActor)