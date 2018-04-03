package engine

import (
	"strconv"
	"strings"
	"path/filepath"
)

// ----------------------------------------------------
// Component manager
// ----------------------------------------------------

type ComponentManager interface {
	Ensure()
	Clean()
	ComponentPath(component Component)
	RepositoryPath()
}

func createComponentManager(ctx context) (ComponentManager, error) {
	cm := componentManager{ctx: ctx, base: ctx.home + string(filepath.Separator) + "components"}

	// Process stacks
	ctx.env.GetStackDescriptions()

	return cm, nil
}

type componentManager struct {
	ctx        context
	base       string
	components []Component
}

func (componentManager) Ensure() {
	panic("implement me")
}

func (componentManager) Clean() {
	panic("implement me")
}

func (componentManager) ComponentPath(component Component) {
	panic("implement me")
}

func (componentManager) RepositoryPath() {
	panic("implement me")
}

// ----------------------------------------------------
// Component
// ----------------------------------------------------

type Component interface {
	Id() string
	Version() Version
	Fetch(version Version, location string)
}

type TaskComponent interface {
	Component
}

type StackComponent interface {
	Component
}

// ----------------------------------------------------
// Version
// ----------------------------------------------------

type Version interface {
	Major() int
	Minor() int
	Micro() int
	AsString() string
}

type version struct {
	major int
	minor int
	micro int
	full  string
}

func (v version) Major() int {
	return v.major
}

func (v version) Minor() int {
	return v.minor
}

func (v version) Micro() int {
	return v.micro
}

func (v version) AsString() string {
	return v.full
}

func CreateVersion(full string) (Version, error) {
	result := version{full: full}
	split := strings.Split(full, ".")
	if len(split) > 0 {
		major, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		result.major = int(major)
	}
	if len(split) > 1 {
		minor, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		result.minor = int(minor)
	}
	if len(split) > 2 {
		minor, err := strconv.Atoi(split[2])
		if err != nil {
			return nil, err
		}
		result.micro = int(minor)
	}
	return result, nil
}
