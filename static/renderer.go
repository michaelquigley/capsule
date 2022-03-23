package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
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
		return nil, errors.Wrapf(err, "error loading procedural def from '%v' (%v)", path, err)
	}
	return def, nil
}

type Renderer interface {
	Render(m *capsule.Model, n *Node) (string, error)
}

func RegisterRenderer(id string, fs cf.FlexibleSetter) {
	if rendererRegistry == nil {
		rendererRegistry = make(map[string]cf.FlexibleSetter)
	}
	rendererRegistry[id] = fs
}

const RendererFeature = "renderer.yaml"

var rendererRegistry map[string]cf.FlexibleSetter
