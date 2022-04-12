package static

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"html/template"
)

func funcMap(m *capsule.Model) template.FuncMap {
	return template.FuncMap{
		"node": func(n *capsule.Node) *Node {
			return newNode(n, m)
		},
		"unescape": func(v interface{}) template.HTML {
			return template.HTML(fmt.Sprintf("%v", v))
		},
	}
}
