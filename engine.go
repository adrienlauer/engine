package engine

import (
	"log"
)

//go:generate go run ./generate/generate.go

const (
	DefaultHomeLocation = "/var/lib/lagoon"
)

type context struct {
	logger *log.Logger
	home   string

	env *environmentDef
	cm  *ComponentManager
}

type Engine interface {
	Environment() Environment
	ComponentManager() ComponentManager
}

// Create creates an environment descriptor based on the provider location.
//
// The location can be an URL over http or https or even a file system location.
func Create(logger *log.Logger, location string) (engine Engine, err error) {
	ctx := context{logger: logger, home: DefaultHomeLocation}
	env, err := parseDescriptor(ctx, location)
	if err != nil {
		return nil, err
	}
	ctx.env = &env

	cm, err := createComponentManager(ctx)
	if err != nil {
		return nil, err
	}
	ctx.cm = &cm

	return &ctx, nil
}

// Retrieve the environment definition
func (ctx context) Environment() Environment {
	return *ctx.env
}

// Return the subsystem managing lagoon components
func (ctx context) ComponentManager() ComponentManager {
	return *ctx.cm
}
