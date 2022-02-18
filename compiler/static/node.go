package static

import (
	"github.com/michaelquigley/capsule"
	"path/filepath"
)

// Node wrapper containing convenience functions for accessing from templates.
//
type Node struct {
	*capsule.Node
}

func newNode(n *capsule.Node) *Node {
	return &Node{n}
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
