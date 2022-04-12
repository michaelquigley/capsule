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

type resources struct {
	tmpl     *template.Template
	template map[string]string
	body     map[string][]Renderer
	statics  []string
}

const renderYaml = "render.yaml"
const staticRoot = "static"
const templatesRoot = "templates"

func loadResources(opt *Options, m *capsule.Model) (*resources, error) {
	r := &resources{}
	if err := r.loadTemplates(opt, m); err != nil {
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
	var tpls []string
	err := fs.WalkDir(os.DirFS(filepath.Join(opt.ResourcePath, templatesRoot)), ".", func(path string, de fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".gohtml" {
			tpls = append(tpls, filepath.ToSlash(filepath.Join(opt.ResourcePath, templatesRoot, path)))
		}
		return nil
	})
	if err != nil {
		return errors.Wrap(err, "error loading templates")
	}
	if tmpl, err := template.New("").Funcs(funcMap(m)).ParseFiles(tpls...); err == nil {
		r.tmpl = tmpl
	} else {
		return errors.Wrap(err, "error parsing templates")
	}
	return nil
}

func (r *resources) loadRenderYaml(opt *Options) error {
	renderYamlPath := filepath.Join(opt.ResourcePath, renderYaml)
	_, err := os.Stat(renderYamlPath)
	if os.IsNotExist(err) {
		logrus.Warnf("no %v loaded", renderYamlPath)
		return nil
	}
	if err != nil {
		return err
	}
	if def, err := LoadRenderDef(renderYamlPath); err == nil {
		for _, entry := range def.Render {
			if entry.Template != "" {
				if r.template == nil {
					r.template = make(map[string]string)
				}
				r.template[entry.Path] = entry.Template
			}
			if len(entry.Body) > 0 {
				if r.body == nil {
					r.body = make(map[string][]Renderer)
				}
				var renderers []Renderer
				for _, v := range entry.Body {
					renderers = append(renderers, v.(Renderer))
				}
				r.body[entry.Path] = renderers
			}
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
		if _, err := CopyFile(srcPath, dstPath); err != nil {
			return nil, err
		}
		dstPaths = append(dstPaths, static)
		logrus.Infof("=> '%v'", dstPath)
	}
	return dstPaths, nil
}
