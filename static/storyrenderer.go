package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"os"
	"path/filepath"
)

func init() {
	RegisterRenderer("story", func(v interface{}, opt *cf.Options) (interface{}, error) {
		cfg := DefaultStoryRendererConfig()
		if data, ok := v.(map[string]interface{}); ok {
			if err := cf.Bind(cfg, data, opt); err == nil {
				return &StoryRenderer{cfg}, nil
			} else {
				return nil, err
			}
		} else {
			return nil, errors.Errorf("invalid configuration data for story renderer (%v)", v)
		}
	})
}

type StoryRendererConfig struct {
	Template string
	Filename string
}

func DefaultStoryRendererConfig() *StoryRendererConfig {
	return &StoryRendererConfig{
		Template: "node",
		Filename: "index.html",
	}
}

type StoryRenderer struct {
	cfg *StoryRendererConfig
}

func (sr *StoryRenderer) Render(opt *Options, m *capsule.Model, n *capsule.Node, t *template.Template) ([]string, error) {
	var dstPaths []string

	renderPath := filepath.ToSlash(filepath.Join(opt.BuildPath, n.FullPath(), sr.cfg.Filename))
	if err := os.MkdirAll(filepath.Dir(renderPath), os.ModePerm); err != nil {
		return nil, err
	}
	f, err := os.Create(renderPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	if err := t.ExecuteTemplate(f, sr.cfg.Template, n); err != nil {
		return nil, err
	}
	logrus.Infof("=> '%v'", renderPath)
	dstPaths = append(dstPaths, filepath.ToSlash(filepath.Join(n.FullPath(), sr.cfg.Filename)))

	return dstPaths, nil
}
