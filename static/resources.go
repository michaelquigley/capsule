package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/util"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type resources struct {
	t        *template.Template
	visitors map[string][]Visitor
	body     map[string][]Renderer
	statics  []string
}

const renderYaml = "render.yaml"
const staticRoot = "static"
const templatesRoot = "templates"
const visitorYaml = "visitor.yaml"

func loadResources(opt *Options, m *capsule.Model) (*resources, error) {
	r := &resources{}
	if err := r.loadTemplates(opt, m); err != nil {
		return nil, err
	}
	if err := r.loadVisitorYaml(opt); err != nil {
		return nil, err
	}
	if err := r.loadRenderYaml(opt); err != nil {
		return nil, err
	}
	if err := r.loadStatic(opt); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *resources) loadTemplates(opt *Options, m *capsule.Model) error {
	t := template.New("").Funcs(funcMap(m))
	err := fs.WalkDir(os.DirFS(filepath.Join(opt.ResourcePath, templatesRoot)), ".", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".gohtml" {
			name := strings.TrimSuffix(path, filepath.Ext(path))
			data, err := os.ReadFile(filepath.Join(opt.ResourcePath, templatesRoot, path))
			if err != nil {
				return errors.Wrapf(err, "error reading template '%v'", path)
			}
			t, err = t.New(name).Parse(string(data))
			if err != nil {
				return errors.Wrapf(err, "error parsing template '%v'", path)
			}
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "error loading templates")
	}
	r.t = t
	return nil
}

func (r *resources) loadVisitorYaml(opt *Options) error {
	visitorYamlPath := filepath.ToSlash(filepath.Join(opt.ResourcePath, visitorYaml))
	_, err := os.Stat(visitorYamlPath)
	if os.IsNotExist(err) {
		logrus.Warnf("no %v loaded", visitorYamlPath)
		return nil
	}
	if err != nil {
		return err
	}
	if def, err := LoadVisitorDef(visitorYamlPath); err == nil {
		for _, visitorDef := range def.Visit {
			if r.visitors == nil {
				r.visitors = make(map[string][]Visitor)
			}
			var visitors []Visitor
			for _, v := range visitorDef.Impl {
				visitors = append(visitors, v.(Visitor))
			}
			r.visitors[visitorDef.Glob] = visitors
		}
	} else {
		return err
	}
	return nil
}

func (r *resources) loadRenderYaml(opt *Options) error {
	renderYamlPath := filepath.ToSlash(filepath.Join(opt.ResourcePath, renderYaml))
	_, err := os.Stat(renderYamlPath)
	if os.IsNotExist(err) {
		logrus.Warnf("no %v loaded", renderYamlPath)
		return nil
	}
	if err != nil {
		return err
	}
	if def, err := LoadRenderDef(renderYamlPath); err == nil {
		for _, renderDef := range def.Render {
			if r.body == nil {
				r.body = make(map[string][]Renderer)
			}
			var renderers []Renderer
			for _, v := range renderDef.Impl {
				renderers = append(renderers, v.(Renderer))
			}
			r.body[renderDef.Glob] = renderers
			logrus.Debugf("loaded %d renderers for '%v'", len(renderers), renderDef.Glob)
		}
	} else {
		return err
	}
	return nil
}

func (r *resources) loadStatic(opt *Options) error {
	err := fs.WalkDir(os.DirFS(filepath.Join(opt.ResourcePath, staticRoot)), ".", func(path string, de fs.DirEntry, err error) error {
		if !de.IsDir() {
			r.statics = append(r.statics, filepath.ToSlash(path))
			logrus.Debugf("loaded => '%v'", path)
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "error loading statics")
	}
	return nil
}

func (r *resources) build(opt *Options) ([]string, error) {
	return r.copyStatic(opt)
}

func (r *resources) copyStatic(opt *Options) ([]string, error) {
	var dstPaths []string
	for _, static := range r.statics {
		srcPath := filepath.ToSlash(filepath.Join(opt.ResourcePath, staticRoot, static))
		dstPath := filepath.ToSlash(filepath.Join(opt.BuildPath, static))
		if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
			return nil, err
		}
		if _, err := util.CopyFile(srcPath, dstPath); err != nil {
			return nil, err
		}
		dstPaths = append(dstPaths, static)
		logrus.Infof("=> '%v'", dstPath)
	}
	return dstPaths, nil
}
