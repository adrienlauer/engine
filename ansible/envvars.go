package ansible

import "github.com/ekara-platform/model"

// EnvVars contains the extra vars to be passed to a playbook
type EnvVars struct {
	Content map[string]string
}

//BuildEnvVars creates and empy EnvVars
func BuildEnvVars() EnvVars {
	r := EnvVars{}
	r.Content = make(map[string]string)
	return r
}

// Add adds the given key and value.
//
// If the key already exists then its content will be overwritten by the by the value
func (ev *EnvVars) Add(key, value string) {
	ev.Content[key] = value
}

// Add the proxy information if any
//
// If proxy info is already present, it will be overwritten or removed if empty proxy values are passed
func (ev *EnvVars) AddProxy(proxy model.Proxy) {
	if proxy.Http != "" {
		ev.Content["http_proxy"] = proxy.Http
	} else {
		delete(ev.Content, "http_proxy")
	}
	if proxy.Https != "" {
		ev.Content["https_proxy"] = proxy.Https
	} else {
		delete(ev.Content, "https_proxy")
	}
	if proxy.NoProxy != "" {
		ev.Content["no_proxy"] = proxy.NoProxy
	} else {
		delete(ev.Content, "no_proxy")
	}
}

// AddBuffer adds the environment variable coming from the given buffer
//
// Only the "Extravars" content of the buffer will be processed.
//
// If any of the buffered keys already exist then its content will be
// overwritten by the by the corresponding buffered value.
func (ev *EnvVars) AddBuffer(b Buffer) {
	for k, v := range b.Envvars {
		ev.Content[k] = v
	}
}
