package static

import (
	"github.com/michaelquigley/capsule"
	"github.com/pkg/errors"
	"path/filepath"
)

// Node wrapper containing convenience functions for accessing from templates.
//
type Node struct {
	*capsule.Node
	Model *capsule.Model
}

func newNode(n *capsule.Node, m *capsule.Model) *Node {
	return &Node{n, m}
}

func (n *Node) Title() string {
	return n.FullPath()
}

func (n *Node) ChildPaths() []string {
	var childPaths []string
	for _, child := range n.Children {
		if rel, err := filepath.Rel(n.FullPath(), child.FullPath()); err == nil {
			childPaths = append(childPaths, rel)
		}
	}
	return childPaths
}

func (n *Node) Timeline() (*capsule.TimelineStructure, error) {
	if v, found := n.Model.Structures["timeline"]; found {
		if ts, ok := v.(*capsule.TimelineStructure); ok {
			return ts, nil
		} else {
			return nil, errors.Errorf("invalid assert in timeline")
		}
	}
	return nil, nil
}

func (n *Node) Rel(o *capsule.Node) (string, error) {
	if rel, err := filepath.Rel(n.FullPath(), o.FullPath()); err == nil {
		return filepath.ToSlash(rel), nil
	} else {
		return "", err
	}
}
