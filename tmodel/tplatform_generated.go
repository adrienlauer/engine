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
// TPlatform is a read only platform
// ----------------------------------------------------
type TPlatform interface {
    //Base returns the base location of the platform
    Base() TBase
    //Distribution returns the distribution used by the platform
    Distribution() TComponent
    //HasComponents returns true if the platform has components
    HasComponents() bool
	
}

// ----------------------------------------------------
// Implementation(s) of TPlatform  
// ----------------------------------------------------

//TPlatformOnPlatformHolder is the struct containing the Platform in order to implement TPlatform  
type TPlatformOnPlatformHolder struct {
	h 	model.Platform
}

//CreateTPlatformForPlatform returns an holder of Platform implementing TPlatform
func CreateTPlatformForPlatform(o model.Platform) TPlatformOnPlatformHolder {
	return TPlatformOnPlatformHolder{
		h: o,
	}
}

//Base returns the base location of the platform
func (r TPlatformOnPlatformHolder) Base() TBase{
	    return CreateTBaseForBase(r.h.Base)
}

//Distribution returns the distribution used by the platform
func (r TPlatformOnPlatformHolder) Distribution() TComponent{
	    return CreateTComponentForDistribution(r.h.Distribution)
}

//HasComponents returns true if the platform has components
func (r TPlatformOnPlatformHolder) HasComponents() bool{
	return len(r.h.Components) > 0
}

