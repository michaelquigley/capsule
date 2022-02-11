package compiler

import "github.com/michaelquigley/capsule"

type Var struct {
	Desc    string
	Default interface{}
}

type Def map[string]Var

type Config map[string]interface{}

type Factory interface {
	New(cfg Config) (Compiler, error)
	Doc() Def
}

type Compiler interface {
	Compile(m *capsule.Model) error
}
