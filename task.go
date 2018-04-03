package engine

type tasks struct {
	values namedMap
}

func createTasks(desc *environmentDef) tasks {
	ret := tasks{namedMap{}}
	for k, v := range desc.Tasks {
		v.name = k
		v.desc = desc
		ret.values[k] = v
	}
	return ret
}

func (e tasks) GetTask(candidate string) (TaskDescription, bool) {
	if v, ok := e.values[candidate]; ok {
		return v.(TaskDescription), ok
	}
	return nil, false
}
