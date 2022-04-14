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
	RegisterVisitor("timelineindex", func(v interface{}, opt *cf.Options) (interface{}, error) {
		cfg := DefaultTimelineIndexVisitorConfig()
		if data, ok := v.(map[string]interface{}); ok {
			if err := cf.Bind(cfg, data, opt); err == nil {
				return &TimelineIndexVisitor{cfg}, nil
			} else {
				return nil, err
			}
		} else {
			return nil, errors.Errorf("invalid configuration data for timeline index (%v)", v)
		}
	})
}

type TimelineIndexVisitorConfig struct {
	Id       string
	Template string
}

func DefaultTimelineIndexVisitorConfig() *TimelineIndexVisitorConfig {
	return &TimelineIndexVisitorConfig{
		Id:       "timeline",
		Template: "renderers/timelineindex",
	}
}

type TimelineIndexVisitor struct {
	cfg *TimelineIndexVisitorConfig
}

func (ti *TimelineIndexVisitor) Visit(m *capsule.Model, n *capsule.Node, tmpl *template.Template) error {
	if v, found := m.Structures[ti.cfg.Id]; found {
		if ts, ok := v.(*capsule.TimelineStructure); ok {
			buf := new(bytes.Buffer)
			if err := tmpl.ExecuteTemplate(buf, ti.cfg.Template, &timelineModel{n, ts}); err != nil {
				return err
			}
			body := n.VString(bodyV) + buf.String()
			n.SetV(bodyV, body)
		} else {
			return errors.Errorf("invalid timeline structure type '%v'", reflect.TypeOf(v).Name())
		}
	} else {
		return errors.Errorf("missing timeline structure '%v'", ti.cfg.Id)
	}
	return nil
}

type timelineModel struct {
	IndexNode *capsule.Node
	Timeline  *capsule.TimelineStructure
}
