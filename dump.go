package capsule

import (
	"fmt"
	"path/filepath"
)

func (self *Model) Dump() string {
	out := self.dumpNode("", self.Root)
	out += self.dumpStructures()
	return out
}

func (self *Model) dumpNode(parentPath string, node *Node) string {
	var nodePath string
	if parentPath != "" {
		nodePath = filepath.ToSlash(filepath.Join(parentPath, node.Path))
	} else {
		nodePath = node.Path
	}

	out := fmt.Sprintf("%v:\n", nodePath)
	for _, feature := range node.Features {
		out += fmt.Sprintf("+-- %v %v\n", feature.Name, feature.Attributes)
	}
	out += "\n"
	for _, child := range node.Children {
		out += self.dumpNode(nodePath, child)
	}
	return out
}

func (self *Model) dumpStructures() string {
	out := ""
	for k, v := range self.Structures {
		out += fmt.Sprintf("[%v] = %v\n", k, v)
	}
	return out
}
