package compiler

import "github.com/michaelquigley/capsule"

type Compiler interface {
	Compile(m *capsule.Model) error
}
