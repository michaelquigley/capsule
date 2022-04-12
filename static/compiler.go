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
	res *resources
}

func New(cfg *Options) *compiler {
	return &compiler{opt: cfg}
}

func (cc *compiler) Compile(m *capsule.Model) error {
	if err := cc.loadResources(m); err != nil {
		return err
	}
	staticPaths, err := cc.copyStatic()
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

	staticNode := newNode(n, m)
	if renderers, err := cc.renderersForNode(staticNode); err == nil {
		for _, renderer := range renderers {
			logrus.Debugf("'%v' => %v", staticNode.FullPath(), reflect.TypeOf(renderer))
			out, rendererPaths, err := renderer.Render(cc.opt, m, staticNode, cc.res.tmpl)
			if err == nil {
				staticNode.Body += out
				dstPaths = append(dstPaths, rendererPaths...)
			} else {
				return nil, err
			}
		}
	} else {
		return nil, err
	}

	if err := cc.res.tmpl.ExecuteTemplate(f, "node", staticNode); err != nil {
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

func (cc *compiler) renderersForNode(n *Node) ([]Renderer, error) {
	if cc.res.body != nil {
		if renderers, found := cc.res.body[n.FullPath()]; found {
			return renderers, nil
		}
	}
	return []Renderer{&StoryRenderer{}, &FeaturesRenderer{}}, nil
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
