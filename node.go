package capsule

import "path/filepath"

type Feature struct {
	Name string
	Type string
}

type Node struct {
	Path      string
	Structure []*Structure
	Features  []*Feature
	Parent    *Node
	Children  []*Node
}

func (n *Node) FullPath() string {
	if n.Parent != nil {
		return filepath.Join(n.Parent.FullPath(), n.Path)
	} else {
		return n.Path
	}
}
