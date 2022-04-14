package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"html/template"
	"os"
	"path/filepath"
)

func init() {
	RegisterVisitor("story", func(v interface{}, opt *cf.Options) (interface{}, error) {
		cfg := DefaultStoryVisitorConfig()
		if data, ok := v.(map[string]interface{}); ok {
			if err := cf.Bind(cfg, data, opt); err == nil {
				return &storyVisitor{cfg}, nil
			} else {
				return nil, err
			}
		} else {
			return nil, errors.Errorf("invalid configuration data for story (%v)", v)
		}
	})
}

type storyVisitorConfig struct {
	Template string
}

func DefaultStoryVisitorConfig() *storyVisitorConfig {
	return &storyVisitorConfig{"renderers/story"}
}

type storyVisitor struct {
	cfg *storyVisitorConfig
}

func (sv *storyVisitor) Visit(m *capsule.Model, n *capsule.Node, t *template.Template) error {
	storyFeatures := n.Features.With(capsule.Attributes{"role": "story", "class": "document"})
	if len(storyFeatures) == 1 {
		logrus.Debugf("visiting '%v'", n.FullPath())
		storyPath := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), storyFeatures[0].Name))
		logrus.Debugf("story path = '%v'", storyPath)

		storySrc, err := os.ReadFile(storyPath)
		if err != nil {
			return err
		}

		var mdBuf bytes.Buffer
		if err := goldmark.Convert(storySrc, &mdBuf); err != nil {
			return err
		}

		buf := new(bytes.Buffer)
		if err := t.ExecuteTemplate(buf, sv.cfg.Template, mdBuf.String()); err != nil {
			return err
		}

		body := n.VString(bodyV) + buf.String()
		n.SetV(bodyV, body)
	}

	logrus.Debugf("no single story to render on '%v'", n.FullPath())
	return nil
}
