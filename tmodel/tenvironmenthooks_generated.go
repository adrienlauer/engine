package engine

import(
	"github.com/ekara-platform/model"
)

//*****************************************************************************
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// ----------------------------------------------------
// TEnvironmentHooks is a read only representation of the hooks associated to an environment
// ----------------------------------------------------
type TEnvironmentHooks interface {
    //HasProvision returns true if the hooks has tasks while provisioning
    HasProvision() bool
    //Provision returns the provisioning tasks
    Provision() THook
    //HasDestroy returns true if the hooks has tasks while destroying
    HasDestroy() bool
    //Destroy returns the destroying tasks
    Destroy() THook
    //HasDeploy returns true if the hooks has tasks while deploying
    HasDeploy() bool
    //Deploy returns the deploying tasks
    Deploy() THook
    //HasUndeploy returns true if the hooks has tasks while undeploying
    HasUndeploy() bool
    //Undeploy returns the undeploying tasks
    Undeploy() THook
	
}

// ----------------------------------------------------
// Implementation(s) of TEnvironmentHooks  
// ----------------------------------------------------

//TEnvironmentHooksOnEnvironmentHooksHolder is the struct containing the EnvironmentHooks in order to implement TEnvironmentHooks  
type TEnvironmentHooksOnEnvironmentHooksHolder struct {
	h 	model.EnvironmentHooks
}

//CreateTEnvironmentHooksForEnvironmentHooks returns an holder of EnvironmentHooks implementing TEnvironmentHooks
func CreateTEnvironmentHooksForEnvironmentHooks(o model.EnvironmentHooks) TEnvironmentHooksOnEnvironmentHooksHolder {
	return TEnvironmentHooksOnEnvironmentHooksHolder{
		h: o,
	}
}

//HasProvision returns true if the hooks has tasks while provisioning
func (r TEnvironmentHooksOnEnvironmentHooksHolder) HasProvision() bool{
	return r.h.Provision.HasTasks()
}

//Provision returns the provisioning tasks
func (r TEnvironmentHooksOnEnvironmentHooksHolder) Provision() THook{
	    return CreateTHookForHook(r.h.Provision)
}

//HasDestroy returns true if the hooks has tasks while destroying
func (r TEnvironmentHooksOnEnvironmentHooksHolder) HasDestroy() bool{
	return r.h.Destroy.HasTasks()
}

//Destroy returns the destroying tasks
func (r TEnvironmentHooksOnEnvironmentHooksHolder) Destroy() THook{
	    return CreateTHookForHook(r.h.Destroy)
}

//HasDeploy returns true if the hooks has tasks while deploying
func (r TEnvironmentHooksOnEnvironmentHooksHolder) HasDeploy() bool{
	return r.h.Deploy.HasTasks()
}

//Deploy returns the deploying tasks
func (r TEnvironmentHooksOnEnvironmentHooksHolder) Deploy() THook{
	    return CreateTHookForHook(r.h.Deploy)
}

//HasUndeploy returns true if the hooks has tasks while undeploying
func (r TEnvironmentHooksOnEnvironmentHooksHolder) HasUndeploy() bool{
	return r.h.Undeploy.HasTasks()
}

//Undeploy returns the undeploying tasks
func (r TEnvironmentHooksOnEnvironmentHooksHolder) Undeploy() THook{
	    return CreateTHookForHook(r.h.Undeploy)
}

