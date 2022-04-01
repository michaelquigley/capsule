package static

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/pkg/errors"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
)

type resources struct {
	tmpl *template.Template
}

func (cc *compiler) loadResources(m *capsule.Model) error {
	cc.res = &resources{}
	if err := cc.loadTemplates(m); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) loadTemplates(m *capsule.Model) error {
	var tpls []string
	err := fs.WalkDir(os.DirFS(filepath.Join(cc.cfg.ResourcePath, TemplatesRoot)), ".", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".gohtml" {
			tpls = append(tpls, filepath.ToSlash(filepath.Join(cc.cfg.ResourcePath, TemplatesRoot, path)))
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "error loading resources")
	}
	if tmpl, err := template.New("").Funcs(cc.funcMap(m)).ParseFiles(tpls...); err == nil {
		cc.res.tmpl = tmpl
	} else {
		return errors.Wrap(err, "error parsing templates")
	}
	return nil
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

const TemplatesRoot = "templates"
