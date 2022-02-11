package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/compiler"
	"github.com/pkg/errors"
)

type staticFactory struct{}

func (f *staticFactory) New(cfg compiler.Config) (compiler.Compiler, error) {
	sc := &staticCompiler{
		build: "build/",
	}
	if v, found := cfg["build"]; found {
		if s, ok := v.(string); ok {
			sc.build = s
		} else {
			return nil, errors.Errorf("invalid 'build' value")
		}
	}
	if v, found := cfg["resources"]; found {
		if s, ok := v.(string); ok {
			sc.resources = s
		} else {
			return nil, errors.Errorf("invalid 'resources' value")
		}
	} else {
		return nil, errors.Errorf("missing 'resources' value")
	}
	return sc, nil
}

func (f *staticFactory) Doc() compiler.Def {
	return compiler.Def{
		"build":     compiler.Var{Desc: "build path", Default: "build/"},
		"resources": compiler.Var{Desc: "resources path", Default: "resources/"},
	}
}

type staticCompiler struct {
	build     string
	resources string
}

func (sc *staticCompiler) Compile(m *capsule.Model) error {
	return nil
}

func init() {
	compiler.Register("static", &staticFactory{})
}
