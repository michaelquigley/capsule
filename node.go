package capsule

import "path/filepath"

type Node struct {
	Path     string
	Features Features
	Parent   *Node
	Children []*Node
}

func (n *Node) FullPath() string {
	if n.Parent != nil {
		return filepath.ToSlash(filepath.Join(n.Parent.FullPath(), n.Path))
	} else {
		return filepath.ToSlash(n.Path)
	}
}
