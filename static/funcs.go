package static

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/pkg/errors"
	"html/template"
	"path/filepath"
)

func funcMap(m *capsule.Model) template.FuncMap {
	return template.FuncMap{
		"childPaths": func(n *capsule.Node) []string {
			return childPaths(n)
		},
		"exported": func(n *capsule.Node) capsule.Features {
			return exportedFeatures(n)
		},
		"rel": func(o, n *capsule.Node) (string, error) {
			return rel(o, n)
		},
		"relPath": func(path string, n *capsule.Node) (string, error) {
			return relPath(path, n)
		},
		"timeline": func() (*capsule.TimelineStructure, error) {
			return timeline(m)
		},
		"title": func(n *capsule.Node) string {
			return title(n)
		},
		"unescape": func(v interface{}) template.HTML {
			return unescape(v)
		},
	}
}

func childPaths(n *capsule.Node) []string {
	var childPaths []string
	for _, child := range n.Children {
		if rel, err := filepath.Rel(n.FullPath(), child.FullPath()); err == nil {
			childPaths = append(childPaths, rel)
		}
	}
	return childPaths
}

func exportedFeatures(n *capsule.Node) capsule.Features {
	return n.Features.NameNotIn([]string{
		capsule.CapsuleFeature,
		capsule.StructureFeature,
	}).Without(capsule.Attributes{
		"role": "story",
	})
}

func rel(o, n *capsule.Node) (string, error) {
	return relPath(o.FullPath(), n)
}

func relPath(path string, n *capsule.Node) (string, error) {
	if rel, err := filepath.Rel(n.FullPath(), path); err == nil {
		return filepath.ToSlash(rel), nil
	} else {
		return "", err
	}
}

func timeline(m *capsule.Model) (*capsule.TimelineStructure, error) {
	if v, found := m.Structures["timeline"]; found {
		if ts, ok := v.(*capsule.TimelineStructure); ok {
			return ts, nil
		} else {
			return nil, errors.Errorf("invalid assert in timeline")
		}
	}
	return nil, nil
}

func title(n *capsule.Node) string {
	return n.FullPath()
}

func unescape(v interface{}) template.HTML {
	return template.HTML(fmt.Sprintf("%v", v))
}
