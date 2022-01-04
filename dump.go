package capsule

import "fmt"

func (self *Model) Dump() string {
	return self.dumpNode(self.Root)
}

func (self *Model) dumpNode(node *Node) string {
	out := fmt.Sprintf("%v:\n", node.Path)
	for _, feature := range node.Features {
		out += fmt.Sprintf("+-- %v %v\n", feature.Name, feature.Attributes.String())
	}
	out += "\n"
	for _, child := range node.Children {
		out += self.dumpNode(child)
	}
	return out
}
