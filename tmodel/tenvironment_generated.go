package engine

import(
	"github.com/ekara-platform/model"
)

//*****************************************************************************
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// TEnvironment is a read only environment
type TEnvironment interface {
    //Name returns the name of the environment
    Name() string
    //Qualifier returns the qualifier of the environment
    Qualifier() string
    //Description returns the description of the environment
    Description() string
    //QualifiedName returns the qualified of the environment
    QualifiedName() string
    //Platform returns the platform used to deploy environment
    Platform() TPlatform
    //HasVars returns true if the environment has defined vars
    HasVars() bool
    //Vars returns the environement vars
    Vars() map[string]interface{}
    //Orchestrator returns the orchestrator managing the environment nodes
    Orchestrator() TOrchestrator
    //HasProviders returns true if the environment has providers
    HasProviders() bool
    //Providers returns the environment providers
    Providers() map[string]TProvider
    //HasNodeSets returns true if the environment has nodes
    HasNodeSets() bool
    //Nodesets returns the environment providers
    Nodesets() map[string]TNodeSet
    //HasStacks returns true if the environment has stacks
    HasStacks() bool
    //Stacks returns the environment stacks
    Stacks() map[string]TStack
    //HasTasks returns true if the environment has tasks
    HasTasks() bool
    //Tasks returns the environment tasks
    Tasks() map[string]TTask
    //HasHooks returns true if the environment has hooks
    HasHooks() bool
    //Hooks returns the environment hooks
    Hooks() TEnvironmentHooks
    //HasProvisionHooks returns true if the environment has hooks while provisioning
    HasProvisionHooks() bool
    //HasDeployHooks returns true if the environment has hooks while deploying
    HasDeployHooks() bool
    //HasUndeployHooks returns true if the environment has hooks while undeploying
    HasUndeployHooks() bool
    //HasDestroyHooks returns true if the environment has hooks while destroying
    HasDestroyHooks() bool
    //HasTemplates returns true if the environment has defined templates
    HasTemplates() bool
    //Templates returns the environment templates
    Templates() []string
	
}

// ----------------------------------------------------
// Implementation(s) of TEnvironment  
// ----------------------------------------------------

//TEnvironmentOnEnvironmentHolder is the struct containing the Environment in order to implement TEnvironment  
type TEnvironmentOnEnvironmentHolder struct {
	h 	model.Environment
}

//CreateTEnvironmentForEnvironment returns an holder of Environment implementing TEnvironment
func CreateTEnvironmentForEnvironment(o model.Environment) TEnvironmentOnEnvironmentHolder {
	return TEnvironmentOnEnvironmentHolder{
		h: o,
	}
}

//Name returns the name of the environment
func (r TEnvironmentOnEnvironmentHolder) Name() string{
	return r.h.Name
}

//Qualifier returns the qualifier of the environment
func (r TEnvironmentOnEnvironmentHolder) Qualifier() string{
	return r.h.Qualifier
}

//Description returns the description of the environment
func (r TEnvironmentOnEnvironmentHolder) Description() string{
	return r.h.Description
}

//QualifiedName returns the qualified of the environment
func (r TEnvironmentOnEnvironmentHolder) QualifiedName() string{
	return r.h.Name + "_" + r.h.Qualifier
}

//Platform returns the platform used to deploy environment
func (r TEnvironmentOnEnvironmentHolder) Platform() TPlatform{
	    return CreateTPlatformForPlatform(*r.h.Ekara)
}

//HasVars returns true if the environment has defined vars
func (r TEnvironmentOnEnvironmentHolder) HasVars() bool{
	return len(r.h.Vars) > 0
}

//Vars returns the environement vars
func (r TEnvironmentOnEnvironmentHolder) Vars() map[string]interface{}{
	return r.h.Vars
}

//Orchestrator returns the orchestrator managing the environment nodes
func (r TEnvironmentOnEnvironmentHolder) Orchestrator() TOrchestrator{
	    return CreateTOrchestratorForOrchestrator(r.h.Orchestrator)
}

//HasProviders returns true if the environment has providers
func (r TEnvironmentOnEnvironmentHolder) HasProviders() bool{
	return len(r.h.Providers) > 0
}

//Providers returns the environment providers
func (r TEnvironmentOnEnvironmentHolder) Providers() map[string]TProvider{
	    result := make(map[string]TProvider)
    for k , val := range r.h.Providers{
        result[k] = CreateTProviderForProvider(val)
    }
    return result

}

//HasNodeSets returns true if the environment has nodes
func (r TEnvironmentOnEnvironmentHolder) HasNodeSets() bool{
	return len(r.h.NodeSets) > 0
}

//Nodesets returns the environment providers
func (r TEnvironmentOnEnvironmentHolder) Nodesets() map[string]TNodeSet{
	    result := make(map[string]TNodeSet)
    for k , val := range r.h.NodeSets{
        result[k] = CreateTNodeSetForNodeSet(val)
    }
    return result

}

//HasStacks returns true if the environment has stacks
func (r TEnvironmentOnEnvironmentHolder) HasStacks() bool{
	return len(r.h.Stacks) > 0
}

//Stacks returns the environment stacks
func (r TEnvironmentOnEnvironmentHolder) Stacks() map[string]TStack{
	    result := make(map[string]TStack)
    for k , val := range r.h.Stacks{
        result[k] = CreateTStackForStack(val)
    }
    return result

}

//HasTasks returns true if the environment has tasks
func (r TEnvironmentOnEnvironmentHolder) HasTasks() bool{
	return len(r.h.Tasks) > 0
}

//Tasks returns the environment tasks
func (r TEnvironmentOnEnvironmentHolder) Tasks() map[string]TTask{
	    result := make(map[string]TTask)
    for k , val := range r.h.Tasks{
        result[k] = CreateTTaskForTask(*val)
    }
    return result

}

//HasHooks returns true if the environment has hooks
func (r TEnvironmentOnEnvironmentHolder) HasHooks() bool{
	return r.h.Hooks.HasTasks()
}

//Hooks returns the environment hooks
func (r TEnvironmentOnEnvironmentHolder) Hooks() TEnvironmentHooks{
	    return CreateTEnvironmentHooksForEnvironmentHooks(r.h.Hooks)
}

//HasProvisionHooks returns true if the environment has hooks while provisioning
func (r TEnvironmentOnEnvironmentHolder) HasProvisionHooks() bool{
	return r.h.Hooks.Provision.HasTasks()
}

//HasDeployHooks returns true if the environment has hooks while deploying
func (r TEnvironmentOnEnvironmentHolder) HasDeployHooks() bool{
	return r.h.Hooks.Deploy.HasTasks()
}

//HasUndeployHooks returns true if the environment has hooks while undeploying
func (r TEnvironmentOnEnvironmentHolder) HasUndeployHooks() bool{
	return r.h.Hooks.Undeploy.HasTasks()
}

//HasDestroyHooks returns true if the environment has hooks while destroying
func (r TEnvironmentOnEnvironmentHolder) HasDestroyHooks() bool{
	return r.h.Hooks.Destroy.HasTasks()
}

//HasTemplates returns true if the environment has defined templates
func (r TEnvironmentOnEnvironmentHolder) HasTemplates() bool{
	return len(r.h.Templates.Content) > 0
}

//Templates returns the environment templates
func (r TEnvironmentOnEnvironmentHolder) Templates() []string{
	return r.h.Templates.Content
}

