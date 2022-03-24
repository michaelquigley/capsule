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

func (sr *StoryRenderer) Render(m *capsule.Model, n *Node, _ *template.Template) (string, error) {
	stories := n.FeaturesWith(capsule.Attributes{"role": "story", "class": "document"})
	if len(stories) == 1 {
		storyPath := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), stories[0].Name))
		logrus.Debugf("story path = '%v'", storyPath)

		storySrc, err := os.ReadFile(storyPath)
		if err != nil {
			return "", err
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(storySrc, &buf); err != nil {
			return "", err
		}
		return buf.String(), nil
	}

	logrus.Debugf("no story to render on '%v'", n.FullPath())
	return "", nil
}
