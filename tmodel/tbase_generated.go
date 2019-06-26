package engine

import(
	"github.com/ekara-platform/model"
)

//*****************************************************************************
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// TBase is a read only base location
type TBase interface {
    //Url returns the url where the base refers
    Url() TUrl
	
}

// ----------------------------------------------------
// Implementation(s) of TBase  
// ----------------------------------------------------

//TBaseOnBaseHolder is the struct containing the Base in order to implement TBase  
type TBaseOnBaseHolder struct {
	h 	model.Base
}

//CreateTBaseForBase returns an holder of Base implementing TBase
func CreateTBaseForBase(o model.Base) TBaseOnBaseHolder {
	return TBaseOnBaseHolder{
		h: o,
	}
}

//Url returns the url where the base refers
func (r TBaseOnBaseHolder) Url() TUrl{
	    return CreateTUrlForEkUrl(r.h.Url)
}

