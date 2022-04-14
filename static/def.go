package static

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
)

type VisitorDef struct {
	Visitors []interface{}
}

func LoadVisitorDef(path string) (*VisitorDef, error) {
	def := &VisitorDef{}

	options := cf.DefaultOptions()
	for k, v := range visitorRegistry {
		options.AddFlexibleSetter(k, v)
	}

	if err := cf.BindYaml(def, path, options); err != nil {
		return nil, errors.Wrapf(err, "error loading visitor definition from '%v' (%v)", path, err)
	}

	return def, nil
}

type RenderDef struct {
	Render []*PathDef
}

type PathDef struct {
	Path     string
	Template string
	Body     []interface{}
}

func LoadRenderDef(path string) (*RenderDef, error) {
	def := &RenderDef{}

	options := cf.DefaultOptions()
	for k, v := range rendererRegistry {
		options.AddFlexibleSetter(k, v)
	}

	if err := cf.BindYaml(def, path, options); err != nil {
		return nil, errors.Wrapf(err, "error loading render definition from '%v' (%v)", path, err)
	}

	return def, nil
}
