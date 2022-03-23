package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"html/template"
	"reflect"
)

func init() {
	RegisterRenderer("timeline/index", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &TimelineIndex{}, nil
	})
}

type TimelineIndex struct{}

func (ti *TimelineIndex) Render(m *capsule.Model, n *Node, tmpl *template.Template) (string, error) {
	if v, found := m.Structures[capsule.TimelineStructureName]; found {
		if ts, ok := v.(*capsule.TimelineStructure); ok {
			return ti.renderTimeline(ts, n, tmpl)
		} else {
			return "", errors.Errorf("invalid timeline structure type '%v'", reflect.TypeOf(v).Name())
		}
	} else {
		return "", errors.Errorf("missing timeline structure")
	}
}

func (ti *TimelineIndex) renderTimeline(ts *capsule.TimelineStructure, n *Node, tmpl *template.Template) (string, error) {
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, "timeline/index", &timelineModel{n, ts}); err == nil {
		return buf.String(), nil
	} else {
		return "", err
	}
}

type timelineModel struct {
	IndexNode *Node
	Timeline  *capsule.TimelineStructure
}
