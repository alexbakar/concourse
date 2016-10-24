// This file was generated by counterfeiter
package execfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/clock"
	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/exec"
	"github.com/concourse/atc/worker"
)

type FakeFactory struct {
	GetStub        func(lager.Logger, exec.StepMetadata, worker.ArtifactName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, int, atc.Params, atc.Version, atc.ResourceTypes, time.Duration, time.Duration) exec.StepFactory
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  worker.ArtifactName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  int
		arg10 atc.Params
		arg11 atc.Version
		arg12 atc.ResourceTypes
		arg13 time.Duration
		arg14 time.Duration
	}
	getReturns struct {
		result1 exec.StepFactory
	}
	PutStub        func(lager.Logger, exec.StepMetadata, worker.Identifier, worker.Metadata, exec.PutDelegate, atc.ResourceConfig, atc.Tags, int, atc.Params, atc.ResourceTypes, time.Duration, time.Duration) exec.StepFactory
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  worker.Identifier
		arg4  worker.Metadata
		arg5  exec.PutDelegate
		arg6  atc.ResourceConfig
		arg7  atc.Tags
		arg8  int
		arg9  atc.Params
		arg10 atc.ResourceTypes
		arg11 time.Duration
		arg12 time.Duration
	}
	putReturns struct {
		result1 exec.StepFactory
	}
	DependentGetStub        func(lager.Logger, exec.StepMetadata, worker.ArtifactName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, int, atc.Params, atc.ResourceTypes, time.Duration, time.Duration) exec.StepFactory
	dependentGetMutex       sync.RWMutex
	dependentGetArgsForCall []struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  worker.ArtifactName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  int
		arg10 atc.Params
		arg11 atc.ResourceTypes
		arg12 time.Duration
		arg13 time.Duration
	}
	dependentGetReturns struct {
		result1 exec.StepFactory
	}
	TaskStub        func(lager.Logger, worker.ArtifactName, worker.Identifier, worker.Metadata, exec.TaskDelegate, exec.Privileged, atc.Tags, int, exec.TaskConfigSource, atc.ResourceTypes, map[string]string, map[string]string, string, clock.Clock, time.Duration, time.Duration) exec.StepFactory
	taskMutex       sync.RWMutex
	taskArgsForCall []struct {
		arg1  lager.Logger
		arg2  worker.ArtifactName
		arg3  worker.Identifier
		arg4  worker.Metadata
		arg5  exec.TaskDelegate
		arg6  exec.Privileged
		arg7  atc.Tags
		arg8  int
		arg9  exec.TaskConfigSource
		arg10 atc.ResourceTypes
		arg11 map[string]string
		arg12 map[string]string
		arg13 string
		arg14 clock.Clock
		arg15 time.Duration
		arg16 time.Duration
	}
	taskReturns struct {
		result1 exec.StepFactory
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFactory) Get(arg1 lager.Logger, arg2 exec.StepMetadata, arg3 worker.ArtifactName, arg4 worker.Identifier, arg5 worker.Metadata, arg6 exec.GetDelegate, arg7 atc.ResourceConfig, arg8 atc.Tags, arg9 int, arg10 atc.Params, arg11 atc.Version, arg12 atc.ResourceTypes, arg13 time.Duration, arg14 time.Duration) exec.StepFactory {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  worker.ArtifactName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  int
		arg10 atc.Params
		arg11 atc.Version
		arg12 atc.ResourceTypes
		arg13 time.Duration
		arg14 time.Duration
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14})
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14)
	} else {
		return fake.getReturns.result1
	}
}

func (fake *FakeFactory) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeFactory) GetArgsForCall(i int) (lager.Logger, exec.StepMetadata, worker.ArtifactName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, int, atc.Params, atc.Version, atc.ResourceTypes, time.Duration, time.Duration) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1, fake.getArgsForCall[i].arg2, fake.getArgsForCall[i].arg3, fake.getArgsForCall[i].arg4, fake.getArgsForCall[i].arg5, fake.getArgsForCall[i].arg6, fake.getArgsForCall[i].arg7, fake.getArgsForCall[i].arg8, fake.getArgsForCall[i].arg9, fake.getArgsForCall[i].arg10, fake.getArgsForCall[i].arg11, fake.getArgsForCall[i].arg12, fake.getArgsForCall[i].arg13, fake.getArgsForCall[i].arg14
}

func (fake *FakeFactory) GetReturns(result1 exec.StepFactory) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Put(arg1 lager.Logger, arg2 exec.StepMetadata, arg3 worker.Identifier, arg4 worker.Metadata, arg5 exec.PutDelegate, arg6 atc.ResourceConfig, arg7 atc.Tags, arg8 int, arg9 atc.Params, arg10 atc.ResourceTypes, arg11 time.Duration, arg12 time.Duration) exec.StepFactory {
	fake.putMutex.Lock()
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  worker.Identifier
		arg4  worker.Metadata
		arg5  exec.PutDelegate
		arg6  atc.ResourceConfig
		arg7  atc.Tags
		arg8  int
		arg9  atc.Params
		arg10 atc.ResourceTypes
		arg11 time.Duration
		arg12 time.Duration
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12})
	fake.recordInvocation("Put", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12)
	} else {
		return fake.putReturns.result1
	}
}

func (fake *FakeFactory) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeFactory) PutArgsForCall(i int) (lager.Logger, exec.StepMetadata, worker.Identifier, worker.Metadata, exec.PutDelegate, atc.ResourceConfig, atc.Tags, int, atc.Params, atc.ResourceTypes, time.Duration, time.Duration) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return fake.putArgsForCall[i].arg1, fake.putArgsForCall[i].arg2, fake.putArgsForCall[i].arg3, fake.putArgsForCall[i].arg4, fake.putArgsForCall[i].arg5, fake.putArgsForCall[i].arg6, fake.putArgsForCall[i].arg7, fake.putArgsForCall[i].arg8, fake.putArgsForCall[i].arg9, fake.putArgsForCall[i].arg10, fake.putArgsForCall[i].arg11, fake.putArgsForCall[i].arg12
}

func (fake *FakeFactory) PutReturns(result1 exec.StepFactory) {
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) DependentGet(arg1 lager.Logger, arg2 exec.StepMetadata, arg3 worker.ArtifactName, arg4 worker.Identifier, arg5 worker.Metadata, arg6 exec.GetDelegate, arg7 atc.ResourceConfig, arg8 atc.Tags, arg9 int, arg10 atc.Params, arg11 atc.ResourceTypes, arg12 time.Duration, arg13 time.Duration) exec.StepFactory {
	fake.dependentGetMutex.Lock()
	fake.dependentGetArgsForCall = append(fake.dependentGetArgsForCall, struct {
		arg1  lager.Logger
		arg2  exec.StepMetadata
		arg3  worker.ArtifactName
		arg4  worker.Identifier
		arg5  worker.Metadata
		arg6  exec.GetDelegate
		arg7  atc.ResourceConfig
		arg8  atc.Tags
		arg9  int
		arg10 atc.Params
		arg11 atc.ResourceTypes
		arg12 time.Duration
		arg13 time.Duration
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	fake.recordInvocation("DependentGet", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13})
	fake.dependentGetMutex.Unlock()
	if fake.DependentGetStub != nil {
		return fake.DependentGetStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13)
	} else {
		return fake.dependentGetReturns.result1
	}
}

func (fake *FakeFactory) DependentGetCallCount() int {
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	return len(fake.dependentGetArgsForCall)
}

func (fake *FakeFactory) DependentGetArgsForCall(i int) (lager.Logger, exec.StepMetadata, worker.ArtifactName, worker.Identifier, worker.Metadata, exec.GetDelegate, atc.ResourceConfig, atc.Tags, int, atc.Params, atc.ResourceTypes, time.Duration, time.Duration) {
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	return fake.dependentGetArgsForCall[i].arg1, fake.dependentGetArgsForCall[i].arg2, fake.dependentGetArgsForCall[i].arg3, fake.dependentGetArgsForCall[i].arg4, fake.dependentGetArgsForCall[i].arg5, fake.dependentGetArgsForCall[i].arg6, fake.dependentGetArgsForCall[i].arg7, fake.dependentGetArgsForCall[i].arg8, fake.dependentGetArgsForCall[i].arg9, fake.dependentGetArgsForCall[i].arg10, fake.dependentGetArgsForCall[i].arg11, fake.dependentGetArgsForCall[i].arg12, fake.dependentGetArgsForCall[i].arg13
}

func (fake *FakeFactory) DependentGetReturns(result1 exec.StepFactory) {
	fake.DependentGetStub = nil
	fake.dependentGetReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Task(arg1 lager.Logger, arg2 worker.ArtifactName, arg3 worker.Identifier, arg4 worker.Metadata, arg5 exec.TaskDelegate, arg6 exec.Privileged, arg7 atc.Tags, arg8 int, arg9 exec.TaskConfigSource, arg10 atc.ResourceTypes, arg11 map[string]string, arg12 map[string]string, arg13 string, arg14 clock.Clock, arg15 time.Duration, arg16 time.Duration) exec.StepFactory {
	fake.taskMutex.Lock()
	fake.taskArgsForCall = append(fake.taskArgsForCall, struct {
		arg1  lager.Logger
		arg2  worker.ArtifactName
		arg3  worker.Identifier
		arg4  worker.Metadata
		arg5  exec.TaskDelegate
		arg6  exec.Privileged
		arg7  atc.Tags
		arg8  int
		arg9  exec.TaskConfigSource
		arg10 atc.ResourceTypes
		arg11 map[string]string
		arg12 map[string]string
		arg13 string
		arg14 clock.Clock
		arg15 time.Duration
		arg16 time.Duration
	}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16})
	fake.recordInvocation("Task", []interface{}{arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16})
	fake.taskMutex.Unlock()
	if fake.TaskStub != nil {
		return fake.TaskStub(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11, arg12, arg13, arg14, arg15, arg16)
	} else {
		return fake.taskReturns.result1
	}
}

func (fake *FakeFactory) TaskCallCount() int {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return len(fake.taskArgsForCall)
}

func (fake *FakeFactory) TaskArgsForCall(i int) (lager.Logger, worker.ArtifactName, worker.Identifier, worker.Metadata, exec.TaskDelegate, exec.Privileged, atc.Tags, int, exec.TaskConfigSource, atc.ResourceTypes, map[string]string, map[string]string, string, clock.Clock, time.Duration, time.Duration) {
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return fake.taskArgsForCall[i].arg1, fake.taskArgsForCall[i].arg2, fake.taskArgsForCall[i].arg3, fake.taskArgsForCall[i].arg4, fake.taskArgsForCall[i].arg5, fake.taskArgsForCall[i].arg6, fake.taskArgsForCall[i].arg7, fake.taskArgsForCall[i].arg8, fake.taskArgsForCall[i].arg9, fake.taskArgsForCall[i].arg10, fake.taskArgsForCall[i].arg11, fake.taskArgsForCall[i].arg12, fake.taskArgsForCall[i].arg13, fake.taskArgsForCall[i].arg14, fake.taskArgsForCall[i].arg15, fake.taskArgsForCall[i].arg16
}

func (fake *FakeFactory) TaskReturns(result1 exec.StepFactory) {
	fake.TaskStub = nil
	fake.taskReturns = struct {
		result1 exec.StepFactory
	}{result1}
}

func (fake *FakeFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.dependentGetMutex.RLock()
	defer fake.dependentGetMutex.RUnlock()
	fake.taskMutex.RLock()
	defer fake.taskMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeFactory) recordInvocation(key string, args []interface{}) {
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

var _ exec.Factory = new(FakeFactory)
