package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"reflect"
)

func init() {
	RegisterRenderer("timeline/index", func(v interface{}, opt *cf.Options) (interface{}, error) {
		cfg := DefaultTimelineIndexConfig()
		if data, ok := v.(map[string]interface{}); ok {
			if err := cf.Bind(cfg, data, opt); err == nil {
				return &TimelineIndex{cfg}, nil
			} else {
				logrus.Error(err)
				return nil, err
			}
		} else {
			return nil, errors.Errorf("invalid configuration data for timeline index (%v)", v)
		}
	})
}

type TimelineIndexConfig struct {
	Id       string
	Template string
}

func DefaultTimelineIndexConfig() *TimelineIndexConfig {
	return &TimelineIndexConfig{
		Id:       "timeline",
		Template: "renderers/timelineindex",
	}
}

type TimelineIndex struct {
	cfg *TimelineIndexConfig
}

func (ti *TimelineIndex) Render(_ *Options, m *capsule.Model, n *capsule.Node, tmpl *template.Template) (*RenderResult, error) {
	if v, found := m.Structures[ti.cfg.Id]; found {
		if ts, ok := v.(*capsule.TimelineStructure); ok {
			return ti.renderTimeline(ts, n, tmpl)
		} else {
			return nil, errors.Errorf("invalid timeline structure type '%v'", reflect.TypeOf(v).Name())
		}
	} else {
		return nil, errors.Errorf("missing timeline structure '%v'", ti.cfg.Id)
	}
}

func (ti *TimelineIndex) renderTimeline(ts *capsule.TimelineStructure, n *capsule.Node, tmpl *template.Template) (*RenderResult, error) {
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, ti.cfg.Template, &timelineModel{n, ts}); err == nil {
		return &RenderResult{Body: buf.String()}, nil
	} else {
		return nil, err
	}
}

type timelineModel struct {
	IndexNode *capsule.Node
	Timeline  *capsule.TimelineStructure
}
