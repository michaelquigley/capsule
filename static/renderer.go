package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"html/template"
)

type Renderer interface {
	Render(cfg *Options, m *capsule.Model, n *capsule.Node, tmpl *template.Template) (string, []string, error)
}

func RegisterRenderer(id string, fs cf.FlexibleSetter) {
	if rendererRegistry == nil {
		rendererRegistry = make(map[string]cf.FlexibleSetter)
	}
	rendererRegistry[id] = fs
}

var rendererRegistry map[string]cf.FlexibleSetter
