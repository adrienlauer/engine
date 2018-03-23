package engine

import (
	"log"

	"gopkg.in/yaml.v2"
)

//go:generate go run ./generate/generate.go

type holder struct {
	// Global state
	logger *log.Logger
	home   string

	// Subsystems
	cm *ComponentManager

	// EnvironmentDef info
	location string
	env      *environmentDef
}

type Engine interface {
	Environment() Environment
	ComponentManager() ComponentManager

	Content() ([]byte, error)
}

// Create creates an environment descriptor based on the provider location.
//
// The location can be an URL over http or https or even a file system location.
func Create(logger *log.Logger, location string) (engine Engine, err error) {
	h := holder{logger: logger, location: location, home: DefaultHomeLocation}
	desc, err := parseDescriptor(h, location)
	if err != nil {
		return nil, err
	}
	h.env = &desc
	return &h, nil
}

// Create creates an environment descriptor based on the provider serialized
// content.
//
// The serialized content is typically a fully resolved descriptor without any import left.
func CreateFromContent(logger *log.Logger, content []byte) (engine Engine, err error) {
	h := holder{logger: logger, location: ""}
	desc, err := parseContent(h, content)
	if err != nil {
		return nil, err
	}
	h.env = &desc
	return &h, nil
}

// Retrieve the environment definition
func (h *holder) Environment() Environment {
	return h.env
}

// Return the subsystem managing lagoon components
func (h *holder) ComponentManager() ComponentManager {
	return h.cm
}

// GetContent serializes the content on the environment descriptor
func (h *holder) Content() ([]byte, error) {
	b, e := yaml.Marshal(h.env)
	return b, e
}
