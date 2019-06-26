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
// TStackHooks is a read only representation of the hooks associated to a stack
// ----------------------------------------------------
type TStackHooks interface {
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
// Implementation(s) of TStackHooks  
// ----------------------------------------------------

//TStackHooksOnStackHookHolder is the struct containing the StackHook in order to implement TStackHooks  
type TStackHooksOnStackHookHolder struct {
	h 	model.StackHook
}

//CreateTStackHooksForStackHook returns an holder of StackHook implementing TStackHooks
func CreateTStackHooksForStackHook(o model.StackHook) TStackHooksOnStackHookHolder {
	return TStackHooksOnStackHookHolder{
		h: o,
	}
}

//HasDeploy returns true if the hooks has tasks while deploying
func (r TStackHooksOnStackHookHolder) HasDeploy() bool{
	return r.h.Deploy.HasTasks()
}

//Deploy returns the deploying tasks
func (r TStackHooksOnStackHookHolder) Deploy() THook{
	    return CreateTHookForHook(r.h.Deploy)
}

//HasUndeploy returns true if the hooks has tasks while undeploying
func (r TStackHooksOnStackHookHolder) HasUndeploy() bool{
	return r.h.Undeploy.HasTasks()
}

//Undeploy returns the undeploying tasks
func (r TStackHooksOnStackHookHolder) Undeploy() THook{
	    return CreateTHookForHook(r.h.Undeploy)
}

