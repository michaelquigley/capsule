package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"html/template"
)

type RenderResult struct {
	Body  string
	Paths []string
}

type Renderer interface {
	Render(opt *Options, m *capsule.Model, n *capsule.Node, t *template.Template) (*RenderResult, error)
}

func RegisterRenderer(id string, fs cf.FlexibleSetter) {
	if rendererRegistry == nil {
		rendererRegistry = make(map[string]cf.FlexibleSetter)
	}
	rendererRegistry[id] = fs
}

var rendererRegistry map[string]cf.FlexibleSetter
