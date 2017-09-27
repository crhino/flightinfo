// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"sync"

	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResourceFactoryFactory struct {
	FactoryForStub        func(workerClient worker.Client) resource.ResourceFactory
	factoryForMutex       sync.RWMutex
	factoryForArgsForCall []struct {
		workerClient worker.Client
	}
	factoryForReturns struct {
		result1 resource.ResourceFactory
	}
	factoryForReturnsOnCall map[int]struct {
		result1 resource.ResourceFactory
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceFactoryFactory) FactoryFor(workerClient worker.Client) resource.ResourceFactory {
	fake.factoryForMutex.Lock()
	ret, specificReturn := fake.factoryForReturnsOnCall[len(fake.factoryForArgsForCall)]
	fake.factoryForArgsForCall = append(fake.factoryForArgsForCall, struct {
		workerClient worker.Client
	}{workerClient})
	fake.recordInvocation("FactoryFor", []interface{}{workerClient})
	fake.factoryForMutex.Unlock()
	if fake.FactoryForStub != nil {
		return fake.FactoryForStub(workerClient)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.factoryForReturns.result1
}

func (fake *FakeResourceFactoryFactory) FactoryForCallCount() int {
	fake.factoryForMutex.RLock()
	defer fake.factoryForMutex.RUnlock()
	return len(fake.factoryForArgsForCall)
}

func (fake *FakeResourceFactoryFactory) FactoryForArgsForCall(i int) worker.Client {
	fake.factoryForMutex.RLock()
	defer fake.factoryForMutex.RUnlock()
	return fake.factoryForArgsForCall[i].workerClient
}

func (fake *FakeResourceFactoryFactory) FactoryForReturns(result1 resource.ResourceFactory) {
	fake.FactoryForStub = nil
	fake.factoryForReturns = struct {
		result1 resource.ResourceFactory
	}{result1}
}

func (fake *FakeResourceFactoryFactory) FactoryForReturnsOnCall(i int, result1 resource.ResourceFactory) {
	fake.FactoryForStub = nil
	if fake.factoryForReturnsOnCall == nil {
		fake.factoryForReturnsOnCall = make(map[int]struct {
			result1 resource.ResourceFactory
		})
	}
	fake.factoryForReturnsOnCall[i] = struct {
		result1 resource.ResourceFactory
	}{result1}
}

func (fake *FakeResourceFactoryFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.factoryForMutex.RLock()
	defer fake.factoryForMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceFactoryFactory) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceFactoryFactory = new(FakeResourceFactoryFactory)
