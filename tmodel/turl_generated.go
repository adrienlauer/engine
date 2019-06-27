package tmodel

import (
	"github.com/ekara-platform/model"
)

//*****************************************************************************
//
// This file is autogenerated by "go generate .". Do not modify.
//
//*****************************************************************************

// TURL is a read only ekara url
type TURL interface {
	//String returns the string representation of the whole url
	String() string
	//Scheme returns the url scheme
	Scheme() string
	//Path returns the url path
	Path() string
	//AsFilePath returns the path converted as a file path
	AsFilePath() string
	//Host returns the url host
	Host() string
}

// ----------------------------------------------------
// Implementation(s) of TURL
// ----------------------------------------------------

//TURLOnEkUrlHolder is the struct containing the EkUrl in order to implement TURL
type TURLOnEkUrlHolder struct {
	h model.EkUrl
}

//CreateTURLForEkUrl returns an holder of EkUrl implementing TURL
func CreateTURLForEkUrl(o model.EkUrl) TURLOnEkUrlHolder {
	return TURLOnEkUrlHolder{
		h: o,
	}
}

//String returns the string representation of the whole url
func (r TURLOnEkUrlHolder) String() string {
	return r.h.String()
}

//Scheme returns the url scheme
func (r TURLOnEkUrlHolder) Scheme() string {
	return r.h.Scheme()
}

//Path returns the url path
func (r TURLOnEkUrlHolder) Path() string {
	return r.h.Path()
}

//AsFilePath returns the path converted as a file path
func (r TURLOnEkUrlHolder) AsFilePath() string {
	return r.h.AsFilePath()
}

//Host returns the url host
func (r TURLOnEkUrlHolder) Host() string {
	return r.h.Host()
}