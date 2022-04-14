package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/sirupsen/logrus"
	"io/fs"
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
	r   *resources
}

func New(cfg *Options) *compiler {
	return &compiler{opt: cfg}
}

func (cc *compiler) Compile(m *capsule.Model) error {
	r, err := loadResources(cc.opt, m)
	if err != nil {
		return err
	}
	cc.r = r
	if err := cc.visitNode(m, m.Root); err != nil {
		return err
	}
	staticPaths, err := cc.r.build(cc.opt)
	if err != nil {
		return err
	}
	renderPaths, err := cc.renderNode(m.Root, m)
	if err != nil {
		return err
	}
	if err := cc.clean(append(staticPaths, renderPaths...)); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) visitNode(m *capsule.Model, n *capsule.Node) error {
	for _, v := range cc.r.visitors {
		visitor := v.(Visitor)

		logrus.Debugf("visiting with '%v'", reflect.TypeOf(visitor).Name())
		if err := visitor.Visit(m, n); err != nil {
			return err
		}
	}
	for _, child := range n.Children {
		if err := cc.visitNode(m, child); err != nil {
			return err
		}
	}
	return nil
}

func (cc *compiler) renderNode(n *capsule.Node, m *capsule.Model) ([]string, error) {
	var dstPaths []string

	renderPath := filepath.ToSlash(filepath.Join(cc.opt.BuildPath, n.FullPath(), "index.html"))
	if err := os.MkdirAll(filepath.Dir(renderPath), os.ModePerm); err != nil {
		return nil, err
	}
	f, err := os.Create(renderPath)
	if err != nil {
		return nil, err
	}

	if renderers, err := cc.renderersForNode(n); err == nil {
		for _, renderer := range renderers {
			logrus.Debugf("'%v' => %v", n.FullPath(), reflect.TypeOf(renderer))
			result, err := renderer.Render(cc.opt, m, n, cc.r.t)
			if err == nil {
				if result != nil {
					dstPaths = append(dstPaths, result.Paths...)
				}
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	if err := cc.r.t.ExecuteTemplate(f, "node", n); err != nil {
		return nil, err
	}
	logrus.Infof("=> '%v'", renderPath)
	dstPaths = append(dstPaths, filepath.Join(n.FullPath(), "index.html"))

	for _, cn := range n.Children {
		childPaths, err := cc.renderNode(cn, m)
		if err != nil {
			return nil, err
		}
		dstPaths = append(dstPaths, childPaths...)
	}

	return dstPaths, nil
}

func (cc *compiler) renderersForNode(n *capsule.Node) ([]Renderer, error) {
	if cc.r.body != nil {
		if renderers, found := cc.r.body[n.FullPath()]; found {
			return renderers, nil
		}
	}
	return []Renderer{&FeaturesRenderer{}}, nil
}

func (cc *compiler) clean(buildPaths []string) error {
	index := make(map[string]struct{})
	for _, buildPath := range buildPaths {
		index[filepath.ToSlash(buildPath)] = struct{}{}
	}
	err := fs.WalkDir(os.DirFS(cc.opt.BuildPath), ".", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !de.IsDir() {
			if _, found := index[path]; !found {
				if err := os.Remove(filepath.Join(cc.opt.BuildPath, path)); err == nil {
					logrus.Warnf(path)
				} else {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
