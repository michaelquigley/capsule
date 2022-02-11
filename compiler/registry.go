package compiler

import "github.com/pkg/errors"

func Register(id string, f Factory) {
	if registry == nil {
		registry = make(map[string]Factory)
	}
	registry[id] = f
}

func Get(id string) (Factory, error) {
	if c, found := registry[id]; found {
		return c, nil
	}
	return nil, errors.Errorf("no compiler with id '%v'", id)
}

var registry map[string]Factory
