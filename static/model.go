package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"html/template"
)

type Visitor interface {
	Visit(m *capsule.Model, n *capsule.Node, t *template.Template) error
}

func RegisterVisitor(id string, fs cf.FlexibleSetter) {
	if visitorRegistry == nil {
		visitorRegistry = make(map[string]cf.FlexibleSetter)
	}
	visitorRegistry[id] = fs
}

type Renderer interface {
	Render(opt *Options, m *capsule.Model, n *capsule.Node, t *template.Template) ([]string, error)
}

func RegisterRenderer(id string, fs cf.FlexibleSetter) {
	if rendererRegistry == nil {
		rendererRegistry = make(map[string]cf.FlexibleSetter)
	}
	rendererRegistry[id] = fs
}

var visitorRegistry map[string]cf.FlexibleSetter
var rendererRegistry map[string]cf.FlexibleSetter
