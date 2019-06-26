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
// TProvider is a read only provider
// ----------------------------------------------------
type TProvider interface {
    //Name returns the name of the provider
    Name() string
    //Parameters returns the provider parameters
    Parameters() map[string]interface{}
    //EnvVars returns the provider environment variables
    EnvVars() map[string]string
    //Proxy returns the proxy definition applied to the provider
    Proxy() TProxy
    //Component returns the provider component
    Component() (TComponent, error)
	
}

// ----------------------------------------------------
// Implementation(s) of TProvider  
// ----------------------------------------------------

//TProviderOnProviderHolder is the struct containing the Provider in order to implement TProvider  
type TProviderOnProviderHolder struct {
	h 	model.Provider
}

//CreateTProviderForProvider returns an holder of Provider implementing TProvider
func CreateTProviderForProvider(o model.Provider) TProviderOnProviderHolder {
	return TProviderOnProviderHolder{
		h: o,
	}
}

//Name returns the name of the provider
func (r TProviderOnProviderHolder) Name() string{
	return r.h.Name
}

//Parameters returns the provider parameters
func (r TProviderOnProviderHolder) Parameters() map[string]interface{}{
	return r.h.Parameters
}

//EnvVars returns the provider environment variables
func (r TProviderOnProviderHolder) EnvVars() map[string]string{
	return r.h.EnvVars
}

//Proxy returns the proxy definition applied to the provider
func (r TProviderOnProviderHolder) Proxy() TProxy{
	    return CreateTProxyForProxy(r.h.Proxy)
}

//Component returns the provider component
func (r TProviderOnProviderHolder) Component() (TComponent, error){
	    v, err := r.h.Component()
    return CreateTComponentForComponent(v), err
}

