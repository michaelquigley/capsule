package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/util"
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

func (fr *FeaturesRenderer) Render(opt *Options, m *capsule.Model, n *capsule.Node, _ *template.Template) ([]string, error) {
	dstPaths, err := fr.copyFeatures(opt, m, n)
	if err != nil {
		return nil, err
	}
	return dstPaths, nil
}

func (fr *FeaturesRenderer) copyFeatures(opt *Options, m *capsule.Model, n *capsule.Node) ([]string, error) {
	var dstPaths []string
	exported := exportedFeatures(n)
	for _, ftr := range exported {
		srcPath := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), ftr.Name))
		dstPath := filepath.ToSlash(filepath.Join(opt.BuildPath, n.FullPath(), ftr.Name))
		if _, err := util.CopyFile(srcPath, dstPath); err != nil {
			return nil, err
		}
		logrus.Infof("=> '%v'", dstPath)
		dstPaths = append(dstPaths, filepath.Join(n.FullPath(), ftr.Name))
	}
	return dstPaths, nil
}
