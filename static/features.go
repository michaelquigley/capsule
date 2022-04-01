package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"html/template"
	"path/filepath"
)

func init() {
	RegisterRenderer("features", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &FeaturesRenderer{}, nil
	})
}

type FeaturesRenderer struct{}

func (fr *FeaturesRenderer) Render(opt *Options, m *capsule.Model, n *Node, _ *template.Template) (string, error) {
	if err := fr.copyFeatures(opt, m, n); err != nil {
		return "", err
	}
	return "", nil
}

func (fr *FeaturesRenderer) copyFeatures(opt *Options, m *capsule.Model, n *Node) error {
	exported := n.ExportedFeatures()
	for _, ftr := range exported {
		srcPath := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), ftr.Name))
		dstPath := filepath.ToSlash(filepath.Join(opt.BuildPath, n.FullPath(), ftr.Name))
		if _, err := CopyFile(srcPath, dstPath); err != nil {
			return err
		}
		logrus.Infof("=> '%v'", dstPath)
	}
	return nil
}
