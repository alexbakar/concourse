// Code generated by counterfeiter. DO NOT EDIT.
package credsfakes

import (
	"sync"
	"time"

	"github.com/concourse/concourse/atc/creds"
)

type FakeSecrets struct {
	GetStub        func(string) (interface{}, *time.Time, bool, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 interface{}
		result2 *time.Time
		result3 bool
		result4 error
	}
	getReturnsOnCall map[int]struct {
		result1 interface{}
		result2 *time.Time
		result3 bool
		result4 error
	}
	NewSecretLookupPathsStub        func(string, string) []creds.SecretLookupPath
	newSecretLookupPathsMutex       sync.RWMutex
	newSecretLookupPathsArgsForCall []struct {
		arg1 string
		arg2 string
	}
	newSecretLookupPathsReturns struct {
		result1 []creds.SecretLookupPath
	}
	newSecretLookupPathsReturnsOnCall map[int]struct {
		result1 []creds.SecretLookupPath
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSecrets) Get(arg1 string) (interface{}, *time.Time, bool, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3, fakeReturns.result4
}

func (fake *FakeSecrets) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeSecrets) GetCalls(stub func(string) (interface{}, *time.Time, bool, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeSecrets) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSecrets) GetReturns(result1 interface{}, result2 *time.Time, result3 bool, result4 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 interface{}
		result2 *time.Time
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSecrets) GetReturnsOnCall(i int, result1 interface{}, result2 *time.Time, result3 bool, result4 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 *time.Time
			result3 bool
			result4 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 interface{}
		result2 *time.Time
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeSecrets) NewSecretLookupPaths(arg1 string, arg2 string) []creds.SecretLookupPath {
	fake.newSecretLookupPathsMutex.Lock()
	ret, specificReturn := fake.newSecretLookupPathsReturnsOnCall[len(fake.newSecretLookupPathsArgsForCall)]
	fake.newSecretLookupPathsArgsForCall = append(fake.newSecretLookupPathsArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("NewSecretLookupPaths", []interface{}{arg1, arg2})
	fake.newSecretLookupPathsMutex.Unlock()
	if fake.NewSecretLookupPathsStub != nil {
		return fake.NewSecretLookupPathsStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.newSecretLookupPathsReturns
	return fakeReturns.result1
}

func (fake *FakeSecrets) NewSecretLookupPathsCallCount() int {
	fake.newSecretLookupPathsMutex.RLock()
	defer fake.newSecretLookupPathsMutex.RUnlock()
	return len(fake.newSecretLookupPathsArgsForCall)
}

func (fake *FakeSecrets) NewSecretLookupPathsCalls(stub func(string, string) []creds.SecretLookupPath) {
	fake.newSecretLookupPathsMutex.Lock()
	defer fake.newSecretLookupPathsMutex.Unlock()
	fake.NewSecretLookupPathsStub = stub
}

func (fake *FakeSecrets) NewSecretLookupPathsArgsForCall(i int) (string, string) {
	fake.newSecretLookupPathsMutex.RLock()
	defer fake.newSecretLookupPathsMutex.RUnlock()
	argsForCall := fake.newSecretLookupPathsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSecrets) NewSecretLookupPathsReturns(result1 []creds.SecretLookupPath) {
	fake.newSecretLookupPathsMutex.Lock()
	defer fake.newSecretLookupPathsMutex.Unlock()
	fake.NewSecretLookupPathsStub = nil
	fake.newSecretLookupPathsReturns = struct {
		result1 []creds.SecretLookupPath
	}{result1}
}

func (fake *FakeSecrets) NewSecretLookupPathsReturnsOnCall(i int, result1 []creds.SecretLookupPath) {
	fake.newSecretLookupPathsMutex.Lock()
	defer fake.newSecretLookupPathsMutex.Unlock()
	fake.NewSecretLookupPathsStub = nil
	if fake.newSecretLookupPathsReturnsOnCall == nil {
		fake.newSecretLookupPathsReturnsOnCall = make(map[int]struct {
			result1 []creds.SecretLookupPath
		})
	}
	fake.newSecretLookupPathsReturnsOnCall[i] = struct {
		result1 []creds.SecretLookupPath
	}{result1}
}

func (fake *FakeSecrets) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.newSecretLookupPathsMutex.RLock()
	defer fake.newSecretLookupPathsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSecrets) recordInvocation(key string, args []interface{}) {
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

var _ creds.Secrets = new(FakeSecrets)
