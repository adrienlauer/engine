package engine

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/imdario/mergo"
	"gopkg.in/yaml.v2"
)

type hookDef struct {
	itemDef
	Before []string
	After  []string
}

type platformDef struct {
	itemDef
	Version string
	Proxy struct {
		Http    string
		Https   string
		NoProxy string `yaml:"noProxy"`
	}
}

type providerDef struct {
	namedItemDef
	labelsDef `yaml:",inline"`
	paramsDef `yaml:",inline"`
}

type nodeSetDef struct {
	namedItemDef
	labelsDef `yaml:",inline"`

	Provider struct {
		paramsDef `yaml:",inline"`
		Name string
	}
	Instances int
	Hooks struct {
		Provision hookDef
		Destroy   hookDef
	}
}

type stackDef struct {
	namedItemDef
	labelsDef `yaml:",inline"`

	Repository string
	Version    string
	DeployOn   []string `yaml:"deployOn"`
	Hooks struct {
		Deploy   hookDef
		Undeploy hookDef
	}
}

type taskDef struct {
	namedItemDef
	labelsDef `yaml:",inline"`

	Playbook string
	Cron     string
	RunOn    []string `yaml:"runOn"`
	Hooks struct {
		Execute hookDef
	}
}

type environmentDef struct {
	labelsDef `yaml:",inline"`

	// Global attributes
	Name        string
	Description string
	Version     string

	// Imports
	Imports []string

	// PlatformDef attributes
	Lagoon platformDef

	// Providers
	Providers map[string]providerDef
	providers providers `yaml:"-"`

	// Node sets
	Nodes map[string]nodeSetDef
	nodes nodes `yaml:"-"`

	// Software stacks
	Stacks map[string]stackDef
	stacks stacks `yaml:"-"`

	// Custom tasks
	Tasks map[string]taskDef
	tasks tasks `yaml:"-"`

	// Global hooks
	Hooks struct {
		Init      hookDef
		Provision hookDef
		Deploy    hookDef
		Undeploy  hookDef
		Destroy   hookDef
	}
}

func parseDescriptor(ctx context, location string) (env environmentDef, err error) {
	baseLocation, content, err := readDescriptor(ctx, location)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(content, &env)
	if err != nil {
		return
	}

	err = processImports(ctx, baseLocation, &env)
	if err != nil {
		return
	}

	env.providers = createProviders(&env)
	env.nodes = createNodes(&env)
	env.stacks = createStacks(&env)
	env.tasks = createTasks(&env)

	return
}

func readDescriptor(ctx context, location string) (base string, content []byte, err error) {
	if strings.Index(location, "http") == 0 {
		ctx.logger.Println("Loading URL", location)

		_, err = url.Parse(location)
		if err != nil {
			return
		}

		var response *http.Response
		response, err = http.Get(location)
		if err != nil {
			return
		}
		defer response.Body.Close()

		if response.StatusCode != 200 {
			err = fmt.Errorf("HTTP Error getting the environment descriptor , error code %d", response.StatusCode)
			return
		}

		content, err = ioutil.ReadAll(response.Body)

		i := strings.LastIndex(location, "/")
		base = location[0 : i+1]
	} else {
		ctx.logger.Println("Loading file", location)
		var file *os.File
		file, err = os.Open(location)
		if err != nil {
			return
		}
		defer file.Close()
		content, err = ioutil.ReadAll(file)
		if err != nil {
			return
		}
		var absLocation string
		absLocation, err = filepath.Abs(location)
		if err != nil {
			return
		}
		base = filepath.Dir(absLocation) + string(filepath.Separator)
	}
	return
}

func processImports(ctx context, baseLocation string, env *environmentDef) error {
	if len(env.Imports) > 0 {
		ctx.logger.Println("Processing imports", env.Imports)
		for _, val := range env.Imports {
			ctx.logger.Println("Processing import", baseLocation+val)
			importedDesc, err := parseDescriptor(ctx, baseLocation+val)
			if err != nil {
				return err
			}
			mergo.Merge(env, importedDesc)
		}
		env.Imports = nil
	} else {
		ctx.logger.Println("No import to process")
	}
	return nil
}

type namedMap map[string]interface{}

func mapContains(m namedMap, candidate string) bool {
	_, ok := m[candidate]
	return ok
}

func mapMultipleContains(m namedMap, candidates []string) bool {
	for _, c := range candidates {
		if b := mapContains(m, c); b == false {
			return false
		}
	}
	return true
}

// ----------------------------------------------------
// Item
// ----------------------------------------------------

type Item interface {
	Root() *environmentDef
}

type NamedItem interface {
	Item
	Name() string
}

type itemDef struct {
	root *environmentDef `yaml:"-"`
}

func (i *itemDef) Root() *environmentDef {
	return i.root;
}

type namedItemDef struct {
	itemDef
	name string `yaml:"-"`
}

func (i *namedItemDef) Name() string {
	return i.name
}

// ----------------------------------------------------
// Labels
// ----------------------------------------------------

type Labels interface {
	Contains(...string) bool
	AsStrings() []string
}

type labelsDef struct {
	Labels []string
}

func createLabels(values ...string) Labels {
	ret := labelsDef{make([]string, len(values))}
	copy(ret.Labels, values)
	return ret
}

func (l labelsDef) Contains(candidates ...string) bool {
	for _, l1 := range candidates {
		contains := false
		for _, l2 := range l.Labels {
			if l1 == l2 {
				contains = true
			}
		}
		if !contains {
			return false
		}
	}
	return true
}

func (l labelsDef) AsStrings() []string {
	ret := make([]string, len(l.Labels))
	copy(ret, l.Labels)
	return ret
}

// ----------------------------------------------------
// Parameters
// ----------------------------------------------------

type Parameters interface {
	AsMap() map[string]string
}

type paramsDef struct {
	Params map[string]string
}

func createParameters(p map[string]string) Parameters {
	ret := paramsDef{map[string]string{}}
	for k, v := range p {
		ret.Params[k] = v
	}
	return ret
}

func (p paramsDef) AsMap() map[string]string {
	ret := map[string]string{}
	for k, v := range p.Params {
		ret[k] = v
	}
	return ret
}

// ----------------------------------------------------
// Nodes
// ----------------------------------------------------

type nodes struct {
	values namedMap
}

func createNodes(env *environmentDef) nodes {
	ret := nodes{namedMap{}}
	for k, v := range env.Nodes {
		v.name = k
		v.root = env
		ret.values[k] = v
	}
	return ret
}

func (e nodes) GetNode(candidate string) (NodeDescription, bool) {
	if v, ok := e.values[candidate]; ok {
		return v.(NodeDescription), ok
	}
	return nil, false
}
