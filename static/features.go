package static

import (
	"bytes"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"html/template"
	"path/filepath"
)

func init() {
	RegisterRenderer("features/index", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &FeaturesRenderer{}, nil
	})
}

type FeaturesRenderer struct{}

func (fr *FeaturesRenderer) Render(cfg *Config, m *capsule.Model, n *Node, tmpl *template.Template) (string, error) {
	if err := fr.copyFeatures(cfg, m, n); err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, "features/index", n); err == nil {
		return buf.String(), nil
	} else {
		return "", err
	}
}

func (fr *FeaturesRenderer) copyFeatures(cfg *Config, m *capsule.Model, n *Node) error {
	for _, ftr := range n.Features {
		srcPath := filepath.Join(m.Path, n.FullPath(), ftr.Name)
		dstPath := filepath.Join(cfg.BuildPath, n.FullPath(), ftr.Name)
		if _, err := CopyFile(srcPath, dstPath); err != nil {
			return err
		}
		logrus.Infof("=> '%v'", dstPath)
	}
	return nil
}
