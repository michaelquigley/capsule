package capsule

import (
	"fmt"
	"path/filepath"
)

const capsuleVersion = "v0.1"

type Model struct {
	SrcPath string
	Capsule *Capsule
	Root    *Node
}

type Capsule struct {
	Version string
}

type Feature struct {
	Name       string
	Attributes Attributes
}

type Attributes struct {
	Role  string
	Class string
	Type  string
}

type Node struct {
	Path     string
	Features []*Feature
	Parent   *Node
	Children []*Node
}

func (n *Node) FullPath() string {
	if n.Parent != nil {
		return filepath.Join(n.Parent.FullPath(), n.Path)
	} else {
		return n.Path
	}
}

func (self *Attributes) String() string {
	return fmt.Sprintf("{%v; %v; %v}", self.Role, self.Class, self.Type)
}
