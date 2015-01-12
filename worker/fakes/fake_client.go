// This file was generated by counterfeiter
package fakes

import (
	"sync"

	garden "github.com/cloudfoundry-incubator/garden/api"
	"github.com/concourse/atc/worker"
)

type FakeClient struct {
	CreateStub        func(garden.ContainerSpec) (worker.Container, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 garden.ContainerSpec
	}
	createReturns struct {
		result1 worker.Container
		result2 error
	}
	LookupStub        func(handle string) (worker.Container, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		handle string
	}
	lookupReturns struct {
		result1 worker.Container
		result2 error
	}
}

func (fake *FakeClient) Create(arg1 garden.ContainerSpec) (worker.Container, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 garden.ContainerSpec
	}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeClient) CreateArgsForCall(i int) garden.ContainerSpec {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeClient) CreateReturns(result1 worker.Container, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) Lookup(handle string) (worker.Container, error) {
	fake.lookupMutex.Lock()
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		handle string
	}{handle})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(handle)
	} else {
		return fake.lookupReturns.result1, fake.lookupReturns.result2
	}
}

func (fake *FakeClient) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeClient) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].handle
}

func (fake *FakeClient) LookupReturns(result1 worker.Container, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

var _ worker.Client = new(FakeClient)
