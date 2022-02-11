package capsule

import (
	"fmt"
	"path/filepath"
	"sort"
)

type Model struct {
	Path       string
	Capsule    *Capsule
	Root       *Node
	Structures map[string]interface{}
}

type Capsule struct {
	Version string
}

type Feature struct {
	Name       string
	Attributes Attributes
	Object     interface{}
}

type Attributes map[string]interface{}

type Node struct {
	Path     string
	Features []*Feature
	Parent   *Node
	Children []*Node
}

func (a Attributes) String() string {
	var keys []string
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := "{"
	i := 0
	for _, k := range keys {
		if i > 0 {
			out += " "
		}
		out += fmt.Sprintf("%v:%v", k, a[k])
		i++
	}
	out += "}"
	return out
}

func (n *Node) FullPath() string {
	if n.Parent != nil {
		return filepath.Join(n.Parent.FullPath(), n.Path)
	} else {
		return n.Path
	}
}

const capsuleVersion = "v0.1"
