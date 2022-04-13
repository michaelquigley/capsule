package capsule

import "path/filepath"

type Node struct {
	Path     string
	Features Features
	Parent   *Node
	Children []*Node
	V        map[string]interface{}
}

func (n *Node) FullPath() string {
	if n.Parent != nil {
		return filepath.ToSlash(filepath.Join(n.Parent.FullPath(), n.Path))
	} else {
		return filepath.ToSlash(n.Path)
	}
}

func (n *Node) SetV(key string, v interface{}) {
	if n.V == nil {
		n.V = make(map[string]interface{})
	}
	n.V[key] = v
}

func (n *Node) VString(key string) string {
	if n.V != nil {
		if v, found := n.V[key]; found {
			if s, ok := v.(string); ok {
				return s
			}
		}
	}
	return ""
}
