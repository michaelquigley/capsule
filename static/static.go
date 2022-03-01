package static

import (
	"bytes"
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
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
	nodePath := n.FullPath() + "/index.html"
	renderPath := filepath.ToSlash(filepath.Join(cc.cfg.BuildPath, nodePath))
	if err := os.MkdirAll(filepath.Dir(renderPath), os.ModePerm); err != nil {
		return err
	}
	f, err := os.Create(renderPath)
	if err != nil {
		return err
	}

	staticNode := newNode(n, m)
	story, err := cc.renderStory(staticNode, m)
	if err != nil {
		return err
	}
	staticNode.Story = story

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

func (cc *compiler) renderStory(n *Node, m *capsule.Model) (string, error) {
	stories := n.FeaturesWith(capsule.Attributes{"role": "story", "class": "document"})
	if len(stories) == 1 {
		storyPath := filepath.ToSlash(filepath.Join(m.Path, n.FullPath(), stories[0].Name))
		logrus.Infof("story path = '%v'", storyPath)

		storySrc, err := os.ReadFile(storyPath)
		if err != nil {
			return "", err
		}

		var buf bytes.Buffer
		if err := goldmark.Convert(storySrc, &buf); err != nil {
			return "", err
		}

		return buf.String(), nil
	}
	return "", nil
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
