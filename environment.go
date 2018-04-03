package engine

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
