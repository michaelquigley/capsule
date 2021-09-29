package capsule

import (
	"fmt"
	"reflect"
)

// Dump the logical contents of a Model for debugging and inspection.
//
func Dump(model *Model) string {
	if model == nil {
		return ""
	}
	out := fmt.Sprintf("Model:\n")
	out += fmt.Sprintf("\tCapsule\n\t\tVersion: '%v'\n", model.Capsule.Version)
	out += "\n"

	out += dumpNode(0, model.Root)

	return out
}

func dumpNode(i int, node *Node) string {
	space := level(i)
	out := fmt.Sprintf("%s%s/\n", space, node.Path)
	for _, s := range node.Structure {
		for _, m := range s.Models {
			if m != nil {
				out += fmt.Sprintf("%s\t<%v>\n", space, reflect.TypeOf(m))
			}
		}
	}
	for _, f := range node.Features {
		out += fmt.Sprintf("%s\t%v (%v)\n", space, f.Name, f.Type)
	}
	out += "\n"
	for _, child := range node.Children {
		out += dumpNode(i + 1, child)
	}
	return out
}

func level(i int) string {
	out := ""
	for j := 0; j < i; j++ {
		out += "\t"
	}
	return out
}