package compiler

import "github.com/pkg/errors"

func Register(id string, c Compiler) {
	if registry == nil {
		registry = make(map[string]Compiler)
	}
	registry[id] = c
}

func Get(id string) (Compiler, error) {
	if c, found := registry[id]; found {
		return c, nil
	}
	return nil, errors.Errorf("no compiler '%v'", id)
}

var registry map[string]Compiler
