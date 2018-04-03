package engine

type stacks struct {
	values namedMap
}

func createStacks(desc *environmentDef) stacks {
	ret := stacks{namedMap{}}
	for k, v := range desc.Stacks {
		v.name = k
		v.desc = desc
		ret.values[k] = v
	}
	return ret
}

func (e stacks) GetStack(candidate string) (StackDescription, bool) {
	if v, ok := e.values[candidate]; ok {
		return v.(StackDescription), ok
	}
	return nil, false
}
