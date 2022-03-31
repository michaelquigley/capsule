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
	ftrs, err := fr.copyFeatures(cfg, m, n)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err := tmpl.ExecuteTemplate(buf, "features/index", struct {
		Node     *Node
		Features []*capsule.Feature
	}{n, ftrs}); err == nil {
		return buf.String(), nil
	} else {
		return "", err
	}
}

func (fr *FeaturesRenderer) copyFeatures(cfg *Config, m *capsule.Model, n *Node) (capsule.Features, error) {
	filtered := fr.filterFeatures(n)
	logrus.Infof("filtered %d features", len(filtered))
	for _, ftr := range filtered {
		srcPath := filepath.Join(m.Path, n.FullPath(), ftr.Name)
		dstPath := filepath.Join(cfg.BuildPath, n.FullPath(), ftr.Name)
		if _, err := CopyFile(srcPath, dstPath); err != nil {
			return nil, err
		}
		logrus.Infof("=> '%v'", dstPath)
	}
	return filtered, nil
}

func (fr *FeaturesRenderer) filterFeatures(n *Node) capsule.Features {
	return n.Features.NameNotIn([]string{
		capsule.CapsuleFeature,
		capsule.StructureFeature,
	}).Without(capsule.Attributes{
		"role": "story",
	})
}
