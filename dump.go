package capsule

import (
	"fmt"
	"path/filepath"
)

func (self *Model) Dump() string {
	return self.dumpNode("", self.Root)
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
		out += fmt.Sprintf("+-- %v %v\n", feature.Name, feature.Attributes.String())
	}
	out += "\n"
	for _, child := range node.Children {
		out += self.dumpNode(nodePath, child)
	}
	return out
}
