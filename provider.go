package engine

type providers struct {
	values namedMap
}

func createProviders(desc *environmentDef) providers {
	ret := providers{namedMap{}}
	for k, v := range desc.Providers {
		v.name = k
		v.desc = desc
		ret.values[k] = v
	}
	return ret
}

func (e providers) GetProvider(candidate string) (ProviderDescription, bool) {
	if v, ok := e.values[candidate]; ok {
		return v.(ProviderDescription), ok
	}
	return nil, false
}
