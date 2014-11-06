package ghooks

import (
	"net/http"
	"strings"
)

func On(name string, handler func(req *http.Request)) {
	hooks.Hooks = append(hooks.Hooks, Hook{Name: name, Func: handler})
}

func Emmit(name string, req *http.Request) {
	for _, v := range hooks.Hooks {
		if strings.EqualFold(v.Name, name) {
			v.Func(req)
		}
	}
}
