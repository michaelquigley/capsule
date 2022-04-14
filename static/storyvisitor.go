package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"os"
	"path/filepath"
)

func init() {
	RegisterVisitor("story", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &StoryVisitor{}, nil
	})
}

type StoryVisitor struct{}

func (sv *StoryVisitor) Visit(m *capsule.Model, n *capsule.Node) error {
	storyFeatures := n.Features.With(capsule.Attributes{"role": "story", "class": "document"})
	if len(storyFeatures) == 1 {
		logrus.Infof("visiting '%v'", n.FullPath())
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

		body := n.VString("body") + mdBuf.String()
		n.SetV("body", body)
		logrus.Infof("new body '%v'", body)
	}

	logrus.Debugf("no single story to render on '%v'", n.FullPath())
	return nil
}
