package ghooks

import (
	"strings"
)

func On(name string, handler func(req interface{})) {
	hooks.Hooks = append(hooks.Hooks, Hook{Name: name, Func: handler})
}

func Emmit(name string, req interface{}) {
	for _, v := range hooks.Hooks {
		if strings.EqualFold(v.Name, name) {
			v.Func(req)
		}
	}
}
