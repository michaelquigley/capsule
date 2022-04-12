package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"reflect"
)

type Options struct {
	BuildPath    string
	ResourcePath string
}

type compiler struct {
	opt *Options
	res *resources
}

func New(cfg *Options) *compiler {
	return &compiler{opt: cfg}
}

func (cc *compiler) Compile(m *capsule.Model) error {
	if err := cc.loadResources(m); err != nil {
		return err
	}
	if err := cc.copyStatic(); err != nil {
		return err
	}
	if err := cc.renderNode(m.Root, m); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) renderNode(n *capsule.Node, m *capsule.Model) error {
	renderPath := filepath.ToSlash(filepath.Join(cc.opt.BuildPath, n.FullPath(), "index.html"))
	if err := os.MkdirAll(filepath.Dir(renderPath), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(renderPath)
	if err != nil {
		return err
	}

	staticNode := newNode(n, m)
	if renderers, err := cc.renderersForNode(staticNode); err == nil {
		for _, renderer := range renderers {
			logrus.Debugf("'%v' => %v", staticNode.FullPath(), reflect.TypeOf(renderer))
			if out, err := renderer.Render(cc.opt, m, staticNode, cc.res.tmpl); err == nil {
				staticNode.Body += out
			} else {
				return err
			}
		}
	} else {
		return err
	}

	if err := cc.res.tmpl.ExecuteTemplate(f, "node", staticNode); err != nil {
		return err
	}
	logrus.Infof("=> '%v'", renderPath)

	for _, cn := range n.Children {
		if err := cc.renderNode(cn, m); err != nil {
			return err
		}
	}
	return nil
}

func (cc *compiler) renderersForNode(n *Node) ([]Renderer, error) {
	if cc.res.body != nil {
		if renderers, found := cc.res.body[n.FullPath()]; found {
			return renderers, nil
		}
	}
	return []Renderer{&StoryRenderer{}, &FeaturesRenderer{}}, nil
}
