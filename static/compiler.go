package static

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
)

type Config struct {
	BuildPath    string
	ResourcePath string
}

type compiler struct {
	cfg  *Config
	tmpl *template.Template
}

func New(cfg *Config) *compiler {
	return &compiler{cfg: cfg}
}

func (cc *compiler) Compile(m *capsule.Model) error {
	if err := cc.loadResources(m); err != nil {
		return err
	}
	if err := cc.renderNode(m.Root, m); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) loadResources(m *capsule.Model) error {
	if err := cc.loadTemplates(m); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) loadTemplates(m *capsule.Model) error {
	var tpls []string
	err := fs.WalkDir(os.DirFS(cc.cfg.ResourcePath), ".", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".gohtml" {
			tpls = append(tpls, filepath.ToSlash(filepath.Join(cc.cfg.ResourcePath, path)))
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "error loading resources")
	}
	if tmpl, err := template.New("").Funcs(cc.funcMap(m)).ParseFiles(tpls...); err == nil {
		cc.tmpl = tmpl
	} else {
		return errors.Wrap(err, "error parsing templates")
	}
	return nil
}

func (cc *compiler) renderNode(n *capsule.Node, m *capsule.Model) error {
	renderPath := filepath.ToSlash(filepath.Join(cc.cfg.BuildPath, n.FullPath(), "index.html"))
	if err := os.MkdirAll(filepath.Dir(renderPath), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(renderPath)
	if err != nil {
		return err
	}

	staticNode := newNode(n, m)
	if renderers, err := cc.renderersForNode(m, staticNode); err == nil {
		for _, renderer := range renderers {
			logrus.Debugf("'%v' => %v", staticNode.FullPath(), reflect.TypeOf(renderer))
			if out, err := renderer.Render(cc.cfg, m, staticNode, cc.tmpl); err == nil {
				staticNode.Body += out
			} else {
				return err
			}
		}
	} else {
		return err
	}

	if err := cc.tmpl.ExecuteTemplate(f, "node", staticNode); err != nil {
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

func (cc *compiler) renderersForNode(m *capsule.Model, n *Node) ([]Renderer, error) {
	if ftr := n.Features.Named(RendererFeature); ftr != nil {
		path := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), ftr.Name))
		if def, err := LoadRendererDef(path); err == nil {
			var renderers []Renderer
			for _, renderer := range def.Renderers {
				renderers = append(renderers, renderer.(Renderer))
			}
			return renderers, nil
		} else {
			return nil, err
		}
	} else {
		return []Renderer{&StoryRenderer{}, &FeaturesRenderer{}}, nil
	}
}

func (cc *compiler) funcMap(m *capsule.Model) template.FuncMap {
	return template.FuncMap{
		"node": func(n *capsule.Node) *Node {
			return newNode(n, m)
		},
		"unescape": func(v interface{}) template.HTML {
			return template.HTML(fmt.Sprintf("%v", v))
		},
	}
}
