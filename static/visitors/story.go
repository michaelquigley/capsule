package visitors

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/static"
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
	"os"
	"path/filepath"
)

func init() {
	static.RegisterVisitor("story", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &StoryVisitor{}, nil
	})
}

type StoryVisitor struct{}

func (sv *StoryVisitor) Visit(m *capsule.Model, n *capsule.Node) error {
	storyFeatures := n.Features.With(capsule.Attributes{"role": "story", "class": "document"})
	if len(storyFeatures) == 1 {
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

		n.SetV("body", n.VString("body")+mdBuf.String())
	}

	logrus.Debugf("no single story to render on '%v'", n.FullPath())
	return nil
}
