package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/compiler"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type staticFactory struct{}

func (f *staticFactory) New(cfg compiler.Config) (compiler.Compiler, error) {
	sc := &staticCompiler{
		build:     "build/",
		resources: "resources/",
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
	}
	return sc, nil
}

type staticCompiler struct {
	build     string
	resources string
}

func (sc *staticCompiler) Compile(m *capsule.Model) error {
	logrus.Infof("build = '%v'", sc.build)
	logrus.Infof("resources = '%v'", sc.resources)
	return nil
}

func init() {
	compiler.Register("static", &staticFactory{})
}
