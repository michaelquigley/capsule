package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"html/template"
)

type RendererDef struct {
	Renderers []interface{}
}

func LoadRendererDef(path string) (*RendererDef, error) {
	def := &RendererDef{}

	options := cf.DefaultOptions()
	for k, v := range rendererRegistry {
		options.AddFlexibleSetter(k, v)
	}

	if err := cf.BindYaml(def, path, options); err != nil {
		return nil, errors.Wrapf(err, "error loading renderers from '%v' (%v)", path, err)
	}

	return def, nil
}

type Renderer interface {
	Render(cfg *Config, m *capsule.Model, n *Node, tmpl *template.Template) (string, error)
}

func RegisterRenderer(id string, fs cf.FlexibleSetter) {
	if rendererRegistry == nil {
		rendererRegistry = make(map[string]cf.FlexibleSetter)
	}
	rendererRegistry[id] = fs
}

const RendererFeature = "renderer.yaml"

var rendererRegistry map[string]cf.FlexibleSetter
