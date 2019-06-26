package engine

import (
	"github.com/ekara-platform/model"
)

//*****************************************************************************
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// TOrchestrator is a read only orchestrator
type TOrchestrator interface {
	//Parameters returns the orchestrator parameters
	Parameters() map[string]interface{}
	//EnvVars returns the orchestrator environment variables
	EnvVars() map[string]string
	//Docker returns the orchestrator parameters for docker
	Docker() map[string]interface{}
	//Component returns the orchestrator component
	Component() (TComponent, error)
}

// ----------------------------------------------------
// Implementation(s) of TOrchestrator
// ----------------------------------------------------

//TOrchestratorOnOrchestratorHolder is the struct containing the Orchestrator in order to implement TOrchestrator
type TOrchestratorOnOrchestratorHolder struct {
	h model.Orchestrator
}

//CreateTOrchestratorForOrchestrator returns an holder of Orchestrator implementing TOrchestrator
func CreateTOrchestratorForOrchestrator(o model.Orchestrator) TOrchestratorOnOrchestratorHolder {
	return TOrchestratorOnOrchestratorHolder{
		h: o,
	}
}

//Parameters returns the orchestrator parameters
func (r TOrchestratorOnOrchestratorHolder) Parameters() map[string]interface{} {
	return r.h.Parameters
}

//EnvVars returns the orchestrator environment variables
func (r TOrchestratorOnOrchestratorHolder) EnvVars() map[string]string {
	return r.h.EnvVars
}

//Docker returns the orchestrator parameters for docker
func (r TOrchestratorOnOrchestratorHolder) Docker() map[string]interface{} {
	return r.h.Docker
}

//Component returns the orchestrator component
func (r TOrchestratorOnOrchestratorHolder) Component() (TComponent, error) {
	v, err := r.h.Component()
	return CreateTComponentForComponent(v), err
}
