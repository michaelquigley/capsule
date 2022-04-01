package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"html/template"
	"os"
	"path/filepath"
)

func init() {
	RegisterRenderer("story", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &StoryRenderer{}, nil
	})
}

type StoryRenderer struct{}

func (sr *StoryRenderer) Render(_ *Options, m *capsule.Model, n *Node, tmpl *template.Template) (string, error) {
	stories := n.Features.With(capsule.Attributes{"role": "story", "class": "document"})
	if len(stories) == 1 {
		storyPath := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), stories[0].Name))
		logrus.Debugf("story path = '%v'", storyPath)

		storySrc, err := os.ReadFile(storyPath)
		if err != nil {
			return "", err
		}

		var mdBuf bytes.Buffer
		if err := goldmark.Convert(storySrc, &mdBuf); err != nil {
			return "", err
		}

		var buf bytes.Buffer
		if err := tmpl.ExecuteTemplate(&buf, "renderers/story", mdBuf.String()); err == nil {
			return buf.String(), nil
		}

		return "", err
	}

	logrus.Debugf("no story to render on '%v'", n.FullPath())
	return "", nil
}
