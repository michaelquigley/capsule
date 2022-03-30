package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"html/template"
)

func init() {
	RegisterRenderer("features/index", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &FeaturesRenderer{}, nil
	})
}

type FeaturesRenderer struct{}

func (fr *FeaturesRenderer) Render(_ *capsule.Model, n *Node, tmpl *template.Template) (string, error) {
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, "features/index", n); err == nil {
		return buf.String(), nil
	} else {
		return "", err
	}
}
