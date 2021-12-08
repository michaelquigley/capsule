package capsule

import "path/filepath"

const capsuleVersion = "v1"

type Model struct {
	SrcPath   string
	Capsule   *Capsule
	Root      *Node
}

type Capsule struct {
	Version string
}

type Feature struct {
	Name string
	Type string
}

type Node struct {
	Path      string
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
