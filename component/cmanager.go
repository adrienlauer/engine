package component

import (
	"fmt"
	"log"
	"os"

	"path/filepath"

	"github.com/ekara-platform/engine/component/scm"
	"github.com/ekara-platform/engine/util"
	"github.com/ekara-platform/model"
	"gopkg.in/yaml.v2"
)

const maxFetchIterations = 9

var releaseNothing = func() {
	// Doing nothing and it's fine
}

//ComponentManager represents the common definition of all Component Manager
type (
	ComponentManager interface {
		//RegisterComponent register a new compoment.
		//
		//The registration key of a component is its id.
		//
		//If the component has already been registered it will remain unaffected
		// by a potential registration of a new version of the component.
		RegisterComponent(c model.Component)
		ComponentsPaths() map[string]string
		SaveComponentsPaths(log *log.Logger, dest util.FolderPath) error
		Ensure() error
		Use(cr model.ComponentReferencer) UsableComponent
		Environment() model.Environment
		ContainsFile(name string, in ...model.ComponentReferencer) MatchingPaths
		ContainsDirectory(name string, in ...model.ComponentReferencer) MatchingPaths
	}

	componentDef struct {
		component model.Component
	}

	context struct {
		// Common to all environments
		logger *log.Logger
		data   *model.TemplateContext

		// Local to one environment (in the case multiple environments will be supported)
		directory   string
		components  map[string]componentDef
		paths       map[string]string
		environment *model.Environment
	}

	compNoTemplate struct {
		component model.Component
	}

	// FileMap is used to Marshal the map of downloaded components
	fileMap struct {
		File map[string]string `yaml:"component_path"`
	}
)

func CreateComponentManager(logger *log.Logger, data *model.TemplateContext, baseDir string) ComponentManager {
	return &context{
		logger:      logger,
		environment: nil,
		directory:   filepath.Join(baseDir, "components"),
		components:  map[string]componentDef{},
		paths:       map[string]string{},
		data:        data,
	}
}

func (cm *context) RegisterComponent(c model.Component) {
	if _, ok := cm.components[c.Id]; !ok {
		cm.logger.Println("registering component " + c.Repository.Url.String() + "@" + c.Repository.Ref)
		cm.components[c.Id] = componentDef{
			component: c,
		}
	}
}

func (cm *context) ComponentsPaths() map[string]string {
	return cm.paths
}

func (cm *context) SaveComponentsPaths(log *log.Logger, dest util.FolderPath) error {
	err := cm.Ensure()
	if err != nil {
		return err
	}
	fMap := fileMap{}
	fMap.File = cm.ComponentsPaths()
	b, err := yaml.Marshal(&fMap)
	if err != nil {
		return err
	}
	_, err = util.SaveFile(log, dest, util.ComponentPathsFileName, b)
	if err != nil {
		return err
	}
	return nil
}

func (cm *context) Ensure() error {
	for i := 0; i < maxFetchIterations && cm.isFetchNeeded(); i++ {
		// Fetch all known components
		for cID, c := range cm.components {
			if cm.isComponentFetchNeeded(cID) {
				err := fetchComponent(cm, c.component)
				if err != nil {
					return err
				}

				// Registering the distribution
				if cm.environment != nil && cm.environment.Ekara != nil && cm.environment.Ekara.Distribution.Repository.Url != nil {
					d := model.Component(cm.environment.Ekara.Distribution)
					if _, ok := cm.components[d.Id]; !ok {
						cm.logger.Printf("registering a distribution")
						cm.RegisterComponent(d)
						err := fetchComponent(cm, d)
						if err != nil {
							return err
						}
						continue
					}
				}
			}
		}

		// Register additionally discovered and used components
		if cm.environment != nil && cm.environment.Ekara != nil {
			cm.logger.Printf("registering used components")
			uc, err := cm.environment.Ekara.UsedComponents()
			if err != nil {
				return err
			}
			for _, c := range uc {
				cm.RegisterComponent(c)
			}
		}
	}
	if cm.isFetchNeeded() {
		return fmt.Errorf("not all components have been fetched after %d iterations, check for import loops in descriptors", maxFetchIterations)
	} else {
		return nil
	}
}

func fetchComponent(cm *context, c model.Component) error {

	h, err := scm.GetHandler(cm.logger, cm.directory, c)
	if err != nil {
		return err
	}
	cm.logger.Printf("fetching component %s ", c.Id)
	fComp, err := h()
	if err != nil {
		return err
	}
	// Parse default descriptor
	err = cm.parseComponentDescriptor(fComp)
	if err != nil {
		return err
	}
	cm.logger.Printf("component %s is available in %s", c.Id, fComp.LocalPath)
	// TODO Change cm.paths to a map[string]FetchedComponent
	cm.paths[c.Id] = fComp.LocalPath
	return nil
}

func (cm *context) Environment() model.Environment {
	return *cm.environment
}

func (cm *context) isFetchNeeded() bool {
	for id := range cm.components {
		if cm.isComponentFetchNeeded(id) {
			return true
		}
	}
	return false
}

func (cm *context) isComponentFetchNeeded(id string) bool {
	_, present := cm.paths[id]
	return !present
}

func (cm *context) parseComponentDescriptor(fComp scm.FetchedComponent) error {
	if fComp.HasDescriptor() {
		// Parsing the descriptor
		cEnv, err := model.CreateEnvironment(fComp.DescriptorUrl, cm.data)
		if err != nil {
			return err
		}

		// Merge the resulting environment into the global one
		if cm.environment == nil {
			cm.environment = cEnv
		} else {
			err = cm.environment.Merge(cEnv)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (cm *context) ContainsFile(name string, in ...model.ComponentReferencer) MatchingPaths {
	return cm.contains(false, name, in...)
}

func (cm *context) ContainsDirectory(name string, in ...model.ComponentReferencer) MatchingPaths {
	return cm.contains(true, name, in...)
}

func (cm *context) contains(isFolder bool, name string, in ...model.ComponentReferencer) MatchingPaths {
	res := MatchingPaths{
		Paths: make([]MatchingPath, 0, 0),
	}
	if len(in) > 0 {
		for _, v := range in {
			uv := cm.Use(v)
			if isFolder {
				if ok, match := uv.ContainsDirectory(name); ok {
					res.Paths = append(res.Paths, match)
				}
			} else {
				if ok, match := uv.ContainsFile(name); ok {
					res.Paths = append(res.Paths, match)
				}
			}
			// TODO add realese if no match
		}
	} else {
		for id, path := range cm.paths {
			dir := filepath.Join(path, name)
			if info, err := os.Stat(dir); err == nil && (isFolder == info.IsDir()) {
				m := mPath{
					comp: cm.Use(
						compNoTemplate{
							cm.components[id].component,
						},
					),
					relativePath: name,
				}
				res.Paths = append(res.Paths, m)
			}
		}
	}
	return res
}

func (cm *context) Use(cr model.ComponentReferencer) UsableComponent {
	if ok, patterns := cr.Templatable(); ok {
		path, err := runTemplate(*cm.data, cm.paths[cr.ComponentName()], patterns, cr)
		if err != nil {
			log.Printf("--> GBE Use err %s", err.Error())
		}
		// No error no path then its has not been templated
		if err == nil && path == "" {
			goto TemplateFalse
		}
		return usable{
			cm:        cm,
			path:      path,
			release:   cleanup(path),
			component: cm.components[cr.ComponentName()],
			templated: true,
		}
	}
TemplateFalse:
	return usable{
		cm:        cm,
		release:   releaseNothing,
		path:      filepath.Join(cm.directory, cr.ComponentName()),
		component: cm.components[cr.ComponentName()],
		templated: false,
	}
}

func cleanup(path string) func() {
	return func() {
		os.RemoveAll(path)
	}
}

//Component returns the referenced component
func (r compNoTemplate) Component() (model.Component, error) {
	return r.component, nil
}

//ComponentName returns the referenced component name
func (r compNoTemplate) ComponentName() string {
	return r.component.Id
}

//Templatable indicates if the component has templates
func (r compNoTemplate) Templatable() (bool, model.Patterns) {
	return false, model.Patterns{}
}
