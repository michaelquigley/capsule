package compiler

import "github.com/michaelquigley/capsule"

type Config map[string]interface{}

type Factory interface {
	New(cfg Config) (Compiler, error)
}

type Compiler interface {
	Compile(m *capsule.Model) error
}
