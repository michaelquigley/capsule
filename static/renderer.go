package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"html/template"
)

type Renderer interface {
	Render(cfg *Options, m *capsule.Model, n *Node, tmpl *template.Template) (string, []string, error)
}

func RegisterRenderer(id string, fs cf.FlexibleSetter) {
	if rendererRegistry == nil {
		rendererRegistry = make(map[string]cf.FlexibleSetter)
	}
	rendererRegistry[id] = fs
}

var rendererRegistry map[string]cf.FlexibleSetter

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
