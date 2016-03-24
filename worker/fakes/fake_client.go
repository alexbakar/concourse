// This file was generated by counterfeiter
package fakes

import (
	"os"
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/worker"
	"github.com/pivotal-golang/lager"
)

type FakeClient struct {
	CreateContainerStub        func(lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes) (worker.Container, error)
	createContainerMutex       sync.RWMutex
	createContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 worker.Identifier
		arg5 worker.Metadata
		arg6 worker.ContainerSpec
		arg7 atc.ResourceTypes
	}
	createContainerReturns struct {
		result1 worker.Container
		result2 error
	}
	FindContainerForIdentifierStub        func(lager.Logger, worker.Identifier) (worker.Container, bool, error)
	findContainerForIdentifierMutex       sync.RWMutex
	findContainerForIdentifierArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Identifier
	}
	findContainerForIdentifierReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	LookupContainerStub        func(lager.Logger, string) (worker.Container, bool, error)
	lookupContainerMutex       sync.RWMutex
	lookupContainerArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupContainerReturns struct {
		result1 worker.Container
		result2 bool
		result3 error
	}
	CreateVolumeStub        func(lager.Logger, worker.VolumeIdentifier, worker.VolumeProperties, bool, time.Duration) (worker.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeIdentifier
		arg3 worker.VolumeProperties
		arg4 bool
		arg5 time.Duration
	}
	createVolumeReturns struct {
		result1 worker.Volume
		result2 error
	}
	ListVolumesStub        func(lager.Logger, worker.VolumeProperties) ([]worker.Volume, error)
	listVolumesMutex       sync.RWMutex
	listVolumesArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.VolumeProperties
	}
	listVolumesReturns struct {
		result1 []worker.Volume
		result2 error
	}
	LookupVolumeStub        func(lager.Logger, string) (worker.Volume, bool, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	lookupVolumeReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	SatisfyingStub        func(worker.WorkerSpec, atc.ResourceTypes) (worker.Worker, error)
	satisfyingMutex       sync.RWMutex
	satisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}
	satisfyingReturns struct {
		result1 worker.Worker
		result2 error
	}
	AllSatisfyingStub        func(worker.WorkerSpec, atc.ResourceTypes) ([]worker.Worker, error)
	allSatisfyingMutex       sync.RWMutex
	allSatisfyingArgsForCall []struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}
	allSatisfyingReturns struct {
		result1 []worker.Worker
		result2 error
	}
	GetWorkerStub        func(workerName string) (worker.Worker, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		workerName string
	}
	getWorkerReturns struct {
		result1 worker.Worker
		result2 error
	}
}

func (fake *FakeClient) CreateContainer(arg1 lager.Logger, arg2 <-chan os.Signal, arg3 worker.ImageFetchingDelegate, arg4 worker.Identifier, arg5 worker.Metadata, arg6 worker.ContainerSpec, arg7 atc.ResourceTypes) (worker.Container, error) {
	fake.createContainerMutex.Lock()
	fake.createContainerArgsForCall = append(fake.createContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 <-chan os.Signal
		arg3 worker.ImageFetchingDelegate
		arg4 worker.Identifier
		arg5 worker.Metadata
		arg6 worker.ContainerSpec
		arg7 atc.ResourceTypes
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7})
	fake.createContainerMutex.Unlock()
	if fake.CreateContainerStub != nil {
		return fake.CreateContainerStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	} else {
		return fake.createContainerReturns.result1, fake.createContainerReturns.result2
	}
}

func (fake *FakeClient) CreateContainerCallCount() int {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return len(fake.createContainerArgsForCall)
}

func (fake *FakeClient) CreateContainerArgsForCall(i int) (lager.Logger, <-chan os.Signal, worker.ImageFetchingDelegate, worker.Identifier, worker.Metadata, worker.ContainerSpec, atc.ResourceTypes) {
	fake.createContainerMutex.RLock()
	defer fake.createContainerMutex.RUnlock()
	return fake.createContainerArgsForCall[i].arg1, fake.createContainerArgsForCall[i].arg2, fake.createContainerArgsForCall[i].arg3, fake.createContainerArgsForCall[i].arg4, fake.createContainerArgsForCall[i].arg5, fake.createContainerArgsForCall[i].arg6, fake.createContainerArgsForCall[i].arg7
}

func (fake *FakeClient) CreateContainerReturns(result1 worker.Container, result2 error) {
	fake.CreateContainerStub = nil
	fake.createContainerReturns = struct {
		result1 worker.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) FindContainerForIdentifier(arg1 lager.Logger, arg2 worker.Identifier) (worker.Container, bool, error) {
	fake.findContainerForIdentifierMutex.Lock()
	fake.findContainerForIdentifierArgsForCall = append(fake.findContainerForIdentifierArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Identifier
	}{arg1, arg2})
	fake.findContainerForIdentifierMutex.Unlock()
	if fake.FindContainerForIdentifierStub != nil {
		return fake.FindContainerForIdentifierStub(arg1, arg2)
	} else {
		return fake.findContainerForIdentifierReturns.result1, fake.findContainerForIdentifierReturns.result2, fake.findContainerForIdentifierReturns.result3
	}
}

func (fake *FakeClient) FindContainerForIdentifierCallCount() int {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return len(fake.findContainerForIdentifierArgsForCall)
}

func (fake *FakeClient) FindContainerForIdentifierArgsForCall(i int) (lager.Logger, worker.Identifier) {
	fake.findContainerForIdentifierMutex.RLock()
	defer fake.findContainerForIdentifierMutex.RUnlock()
	return fake.findContainerForIdentifierArgsForCall[i].arg1, fake.findContainerForIdentifierArgsForCall[i].arg2
}

func (fake *FakeClient) FindContainerForIdentifierReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.FindContainerForIdentifierStub = nil
	fake.findContainerForIdentifierReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) LookupContainer(arg1 lager.Logger, arg2 string) (worker.Container, bool, error) {
	fake.lookupContainerMutex.Lock()
	fake.lookupContainerArgsForCall = append(fake.lookupContainerArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.lookupContainerMutex.Unlock()
	if fake.LookupContainerStub != nil {
		return fake.LookupContainerStub(arg1, arg2)
	} else {
		return fake.lookupContainerReturns.result1, fake.lookupContainerReturns.result2, fake.lookupContainerReturns.result3
	}
}

func (fake *FakeClient) LookupContainerCallCount() int {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return len(fake.lookupContainerArgsForCall)
}

func (fake *FakeClient) LookupContainerArgsForCall(i int) (lager.Logger, string) {
	fake.lookupContainerMutex.RLock()
	defer fake.lookupContainerMutex.RUnlock()
	return fake.lookupContainerArgsForCall[i].arg1, fake.lookupContainerArgsForCall[i].arg2
}

func (fake *FakeClient) LookupContainerReturns(result1 worker.Container, result2 bool, result3 error) {
	fake.LookupContainerStub = nil
	fake.lookupContainerReturns = struct {
		result1 worker.Container
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) CreateVolume(arg1 lager.Logger, arg2 worker.VolumeIdentifier, arg3 worker.VolumeProperties, arg4 bool, arg5 time.Duration) (worker.Volume, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeIdentifier
		arg3 worker.VolumeProperties
		arg4 bool
		arg5 time.Duration
	}{arg1, arg2, arg3, arg4, arg5})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1, arg2, arg3, arg4, arg5)
	} else {
		return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
	}
}

func (fake *FakeClient) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeClient) CreateVolumeArgsForCall(i int) (lager.Logger, worker.VolumeIdentifier, worker.VolumeProperties, bool, time.Duration) {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].arg1, fake.createVolumeArgsForCall[i].arg2, fake.createVolumeArgsForCall[i].arg3, fake.createVolumeArgsForCall[i].arg4, fake.createVolumeArgsForCall[i].arg5
}

func (fake *FakeClient) CreateVolumeReturns(result1 worker.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) ListVolumes(arg1 lager.Logger, arg2 worker.VolumeProperties) ([]worker.Volume, error) {
	fake.listVolumesMutex.Lock()
	fake.listVolumesArgsForCall = append(fake.listVolumesArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.VolumeProperties
	}{arg1, arg2})
	fake.listVolumesMutex.Unlock()
	if fake.ListVolumesStub != nil {
		return fake.ListVolumesStub(arg1, arg2)
	} else {
		return fake.listVolumesReturns.result1, fake.listVolumesReturns.result2
	}
}

func (fake *FakeClient) ListVolumesCallCount() int {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return len(fake.listVolumesArgsForCall)
}

func (fake *FakeClient) ListVolumesArgsForCall(i int) (lager.Logger, worker.VolumeProperties) {
	fake.listVolumesMutex.RLock()
	defer fake.listVolumesMutex.RUnlock()
	return fake.listVolumesArgsForCall[i].arg1, fake.listVolumesArgsForCall[i].arg2
}

func (fake *FakeClient) ListVolumesReturns(result1 []worker.Volume, result2 error) {
	fake.ListVolumesStub = nil
	fake.listVolumesReturns = struct {
		result1 []worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) LookupVolume(arg1 lager.Logger, arg2 string) (worker.Volume, bool, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(arg1, arg2)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2, fake.lookupVolumeReturns.result3
	}
}

func (fake *FakeClient) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeClient) LookupVolumeArgsForCall(i int) (lager.Logger, string) {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].arg1, fake.lookupVolumeArgsForCall[i].arg2
}

func (fake *FakeClient) LookupVolumeReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeClient) Satisfying(arg1 worker.WorkerSpec, arg2 atc.ResourceTypes) (worker.Worker, error) {
	fake.satisfyingMutex.Lock()
	fake.satisfyingArgsForCall = append(fake.satisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}{arg1, arg2})
	fake.satisfyingMutex.Unlock()
	if fake.SatisfyingStub != nil {
		return fake.SatisfyingStub(arg1, arg2)
	} else {
		return fake.satisfyingReturns.result1, fake.satisfyingReturns.result2
	}
}

func (fake *FakeClient) SatisfyingCallCount() int {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return len(fake.satisfyingArgsForCall)
}

func (fake *FakeClient) SatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.ResourceTypes) {
	fake.satisfyingMutex.RLock()
	defer fake.satisfyingMutex.RUnlock()
	return fake.satisfyingArgsForCall[i].arg1, fake.satisfyingArgsForCall[i].arg2
}

func (fake *FakeClient) SatisfyingReturns(result1 worker.Worker, result2 error) {
	fake.SatisfyingStub = nil
	fake.satisfyingReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) AllSatisfying(arg1 worker.WorkerSpec, arg2 atc.ResourceTypes) ([]worker.Worker, error) {
	fake.allSatisfyingMutex.Lock()
	fake.allSatisfyingArgsForCall = append(fake.allSatisfyingArgsForCall, struct {
		arg1 worker.WorkerSpec
		arg2 atc.ResourceTypes
	}{arg1, arg2})
	fake.allSatisfyingMutex.Unlock()
	if fake.AllSatisfyingStub != nil {
		return fake.AllSatisfyingStub(arg1, arg2)
	} else {
		return fake.allSatisfyingReturns.result1, fake.allSatisfyingReturns.result2
	}
}

func (fake *FakeClient) AllSatisfyingCallCount() int {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return len(fake.allSatisfyingArgsForCall)
}

func (fake *FakeClient) AllSatisfyingArgsForCall(i int) (worker.WorkerSpec, atc.ResourceTypes) {
	fake.allSatisfyingMutex.RLock()
	defer fake.allSatisfyingMutex.RUnlock()
	return fake.allSatisfyingArgsForCall[i].arg1, fake.allSatisfyingArgsForCall[i].arg2
}

func (fake *FakeClient) AllSatisfyingReturns(result1 []worker.Worker, result2 error) {
	fake.AllSatisfyingStub = nil
	fake.allSatisfyingReturns = struct {
		result1 []worker.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeClient) GetWorker(workerName string) (worker.Worker, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		workerName string
	}{workerName})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(workerName)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2
	}
}

func (fake *FakeClient) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeClient) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].workerName
}

func (fake *FakeClient) GetWorkerReturns(result1 worker.Worker, result2 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 worker.Worker
		result2 error
	}{result1, result2}
}

var _ worker.Client = new(FakeClient)
