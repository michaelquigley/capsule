package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
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
	if err := cc.loadResources(); err != nil {
		return err
	}
	if err := cc.renderNode(m.Root); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) loadResources() error {
	if err := cc.loadTemplates(); err != nil {
		return err
	}
	return nil
}

func (cc *compiler) loadTemplates() error {
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
	if tmpl, err := template.ParseFiles(tpls...); err == nil {
		cc.tmpl = tmpl
	} else {
		return errors.Wrap(err, "error parsing templates")
	}
	return nil
}

func (cc *compiler) renderNode(n *capsule.Node) error {
	nodePath := n.FullPath() + "/index.html"
	renderPath := filepath.Join(cc.cfg.BuildPath, nodePath)
	if err := os.MkdirAll(filepath.Dir(renderPath), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(renderPath)
	if err != nil {
		return err
	}
	if err := cc.tmpl.ExecuteTemplate(f, "node", n); err != nil {
		return err
	}
	logrus.Infof("=> '%v'", renderPath)
	for _, cn := range n.Children {
		if err := cc.renderNode(cn); err != nil {
			return err
		}
	}
	return nil
}
