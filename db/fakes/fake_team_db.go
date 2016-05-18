// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
)

type FakeTeamDB struct {
	GetAllPipelinesStub        func() ([]db.SavedPipeline, error)
	getAllPipelinesMutex       sync.RWMutex
	getAllPipelinesArgsForCall []struct{}
	getAllPipelinesReturns     struct {
		result1 []db.SavedPipeline
		result2 error
	}
	GetPipelineByNameStub        func(pipelineName string) (db.SavedPipeline, error)
	getPipelineByNameMutex       sync.RWMutex
	getPipelineByNameArgsForCall []struct {
		pipelineName string
	}
	getPipelineByNameReturns struct {
		result1 db.SavedPipeline
		result2 error
	}
	OrderPipelinesStub        func([]string) error
	orderPipelinesMutex       sync.RWMutex
	orderPipelinesArgsForCall []struct {
		arg1 []string
	}
	orderPipelinesReturns struct {
		result1 error
	}
	GetTeamStub        func() (db.SavedTeam, bool, error)
	getTeamMutex       sync.RWMutex
	getTeamArgsForCall []struct{}
	getTeamReturns     struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	UpdateBasicAuthStub        func(basicAuth db.BasicAuth) (db.SavedTeam, error)
	updateBasicAuthMutex       sync.RWMutex
	updateBasicAuthArgsForCall []struct {
		basicAuth db.BasicAuth
	}
	updateBasicAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	UpdateGitHubAuthStub        func(gitHubAuth db.GitHubAuth) (db.SavedTeam, error)
	updateGitHubAuthMutex       sync.RWMutex
	updateGitHubAuthArgsForCall []struct {
		gitHubAuth db.GitHubAuth
	}
	updateGitHubAuthReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	GetConfigStub        func(pipelineName string) (atc.Config, atc.RawConfig, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct {
		pipelineName string
	}
	getConfigReturns struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}
	SaveConfigStub        func(string, atc.Config, db.ConfigVersion, db.PipelinePausedState) (db.SavedPipeline, bool, error)
	saveConfigMutex       sync.RWMutex
	saveConfigArgsForCall []struct {
		arg1 string
		arg2 atc.Config
		arg3 db.ConfigVersion
		arg4 db.PipelinePausedState
	}
	saveConfigReturns struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}
}

func (fake *FakeTeamDB) GetAllPipelines() ([]db.SavedPipeline, error) {
	fake.getAllPipelinesMutex.Lock()
	fake.getAllPipelinesArgsForCall = append(fake.getAllPipelinesArgsForCall, struct{}{})
	fake.getAllPipelinesMutex.Unlock()
	if fake.GetAllPipelinesStub != nil {
		return fake.GetAllPipelinesStub()
	} else {
		return fake.getAllPipelinesReturns.result1, fake.getAllPipelinesReturns.result2
	}
}

func (fake *FakeTeamDB) GetAllPipelinesCallCount() int {
	fake.getAllPipelinesMutex.RLock()
	defer fake.getAllPipelinesMutex.RUnlock()
	return len(fake.getAllPipelinesArgsForCall)
}

func (fake *FakeTeamDB) GetAllPipelinesReturns(result1 []db.SavedPipeline, result2 error) {
	fake.GetAllPipelinesStub = nil
	fake.getAllPipelinesReturns = struct {
		result1 []db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetPipelineByName(pipelineName string) (db.SavedPipeline, error) {
	fake.getPipelineByNameMutex.Lock()
	fake.getPipelineByNameArgsForCall = append(fake.getPipelineByNameArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.getPipelineByNameMutex.Unlock()
	if fake.GetPipelineByNameStub != nil {
		return fake.GetPipelineByNameStub(pipelineName)
	} else {
		return fake.getPipelineByNameReturns.result1, fake.getPipelineByNameReturns.result2
	}
}

func (fake *FakeTeamDB) GetPipelineByNameCallCount() int {
	fake.getPipelineByNameMutex.RLock()
	defer fake.getPipelineByNameMutex.RUnlock()
	return len(fake.getPipelineByNameArgsForCall)
}

func (fake *FakeTeamDB) GetPipelineByNameArgsForCall(i int) string {
	fake.getPipelineByNameMutex.RLock()
	defer fake.getPipelineByNameMutex.RUnlock()
	return fake.getPipelineByNameArgsForCall[i].pipelineName
}

func (fake *FakeTeamDB) GetPipelineByNameReturns(result1 db.SavedPipeline, result2 error) {
	fake.GetPipelineByNameStub = nil
	fake.getPipelineByNameReturns = struct {
		result1 db.SavedPipeline
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) OrderPipelines(arg1 []string) error {
	fake.orderPipelinesMutex.Lock()
	fake.orderPipelinesArgsForCall = append(fake.orderPipelinesArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.orderPipelinesMutex.Unlock()
	if fake.OrderPipelinesStub != nil {
		return fake.OrderPipelinesStub(arg1)
	} else {
		return fake.orderPipelinesReturns.result1
	}
}

func (fake *FakeTeamDB) OrderPipelinesCallCount() int {
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	return len(fake.orderPipelinesArgsForCall)
}

func (fake *FakeTeamDB) OrderPipelinesArgsForCall(i int) []string {
	fake.orderPipelinesMutex.RLock()
	defer fake.orderPipelinesMutex.RUnlock()
	return fake.orderPipelinesArgsForCall[i].arg1
}

func (fake *FakeTeamDB) OrderPipelinesReturns(result1 error) {
	fake.OrderPipelinesStub = nil
	fake.orderPipelinesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTeamDB) GetTeam() (db.SavedTeam, bool, error) {
	fake.getTeamMutex.Lock()
	fake.getTeamArgsForCall = append(fake.getTeamArgsForCall, struct{}{})
	fake.getTeamMutex.Unlock()
	if fake.GetTeamStub != nil {
		return fake.GetTeamStub()
	} else {
		return fake.getTeamReturns.result1, fake.getTeamReturns.result2, fake.getTeamReturns.result3
	}
}

func (fake *FakeTeamDB) GetTeamCallCount() int {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return len(fake.getTeamArgsForCall)
}

func (fake *FakeTeamDB) GetTeamReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamStub = nil
	fake.getTeamReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeTeamDB) UpdateBasicAuth(basicAuth db.BasicAuth) (db.SavedTeam, error) {
	fake.updateBasicAuthMutex.Lock()
	fake.updateBasicAuthArgsForCall = append(fake.updateBasicAuthArgsForCall, struct {
		basicAuth db.BasicAuth
	}{basicAuth})
	fake.updateBasicAuthMutex.Unlock()
	if fake.UpdateBasicAuthStub != nil {
		return fake.UpdateBasicAuthStub(basicAuth)
	} else {
		return fake.updateBasicAuthReturns.result1, fake.updateBasicAuthReturns.result2
	}
}

func (fake *FakeTeamDB) UpdateBasicAuthCallCount() int {
	fake.updateBasicAuthMutex.RLock()
	defer fake.updateBasicAuthMutex.RUnlock()
	return len(fake.updateBasicAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateBasicAuthArgsForCall(i int) db.BasicAuth {
	fake.updateBasicAuthMutex.RLock()
	defer fake.updateBasicAuthMutex.RUnlock()
	return fake.updateBasicAuthArgsForCall[i].basicAuth
}

func (fake *FakeTeamDB) UpdateBasicAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateBasicAuthStub = nil
	fake.updateBasicAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) UpdateGitHubAuth(gitHubAuth db.GitHubAuth) (db.SavedTeam, error) {
	fake.updateGitHubAuthMutex.Lock()
	fake.updateGitHubAuthArgsForCall = append(fake.updateGitHubAuthArgsForCall, struct {
		gitHubAuth db.GitHubAuth
	}{gitHubAuth})
	fake.updateGitHubAuthMutex.Unlock()
	if fake.UpdateGitHubAuthStub != nil {
		return fake.UpdateGitHubAuthStub(gitHubAuth)
	} else {
		return fake.updateGitHubAuthReturns.result1, fake.updateGitHubAuthReturns.result2
	}
}

func (fake *FakeTeamDB) UpdateGitHubAuthCallCount() int {
	fake.updateGitHubAuthMutex.RLock()
	defer fake.updateGitHubAuthMutex.RUnlock()
	return len(fake.updateGitHubAuthArgsForCall)
}

func (fake *FakeTeamDB) UpdateGitHubAuthArgsForCall(i int) db.GitHubAuth {
	fake.updateGitHubAuthMutex.RLock()
	defer fake.updateGitHubAuthMutex.RUnlock()
	return fake.updateGitHubAuthArgsForCall[i].gitHubAuth
}

func (fake *FakeTeamDB) UpdateGitHubAuthReturns(result1 db.SavedTeam, result2 error) {
	fake.UpdateGitHubAuthStub = nil
	fake.updateGitHubAuthReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamDB) GetConfig(pipelineName string) (atc.Config, atc.RawConfig, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct {
		pipelineName string
	}{pipelineName})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub(pipelineName)
	} else {
		return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3, fake.getConfigReturns.result4
	}
}

func (fake *FakeTeamDB) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeTeamDB) GetConfigArgsForCall(i int) string {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return fake.getConfigArgsForCall[i].pipelineName
}

func (fake *FakeTeamDB) GetConfigReturns(result1 atc.Config, result2 atc.RawConfig, result3 db.ConfigVersion, result4 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 atc.RawConfig
		result3 db.ConfigVersion
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeTeamDB) SaveConfig(arg1 string, arg2 atc.Config, arg3 db.ConfigVersion, arg4 db.PipelinePausedState) (db.SavedPipeline, bool, error) {
	fake.saveConfigMutex.Lock()
	fake.saveConfigArgsForCall = append(fake.saveConfigArgsForCall, struct {
		arg1 string
		arg2 atc.Config
		arg3 db.ConfigVersion
		arg4 db.PipelinePausedState
	}{arg1, arg2, arg3, arg4})
	fake.saveConfigMutex.Unlock()
	if fake.SaveConfigStub != nil {
		return fake.SaveConfigStub(arg1, arg2, arg3, arg4)
	} else {
		return fake.saveConfigReturns.result1, fake.saveConfigReturns.result2, fake.saveConfigReturns.result3
	}
}

func (fake *FakeTeamDB) SaveConfigCallCount() int {
	fake.saveConfigMutex.RLock()
	defer fake.saveConfigMutex.RUnlock()
	return len(fake.saveConfigArgsForCall)
}

func (fake *FakeTeamDB) SaveConfigArgsForCall(i int) (string, atc.Config, db.ConfigVersion, db.PipelinePausedState) {
	fake.saveConfigMutex.RLock()
	defer fake.saveConfigMutex.RUnlock()
	return fake.saveConfigArgsForCall[i].arg1, fake.saveConfigArgsForCall[i].arg2, fake.saveConfigArgsForCall[i].arg3, fake.saveConfigArgsForCall[i].arg4
}

func (fake *FakeTeamDB) SaveConfigReturns(result1 db.SavedPipeline, result2 bool, result3 error) {
	fake.SaveConfigStub = nil
	fake.saveConfigReturns = struct {
		result1 db.SavedPipeline
		result2 bool
		result3 error
	}{result1, result2, result3}
}

var _ db.TeamDB = new(FakeTeamDB)
