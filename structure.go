package capsule

import (
	"github.com/pkg/errors"
	"sort"
)

type TimelineStructure struct {
	Nodes []*Node
}

type TimelineStructuralDirective struct{}

func (self *TimelineStructuralDirective) Build(rootPath string, node *Node, prev interface{}) (interface{}, error) {
	prevarr, ok := prev.([]string)
	if !ok {
		return nil, errors.Errorf("invalid previous state")
	}
	nextarr := make([]string, len(prevarr))
	copy(nextarr, prevarr)

	nodearr := self.inventory(node)
	nextarr = append(nextarr, nodearr...)
	sort.Slice(nextarr, func(i, j int) bool {
		return nextarr[i] < nextarr[j]
	})
	return nextarr, nil
}

func (self *TimelineStructuralDirective) inventory(node *Node) []string {
	var paths []string
	for _, cld := range node.Children {
		paths = append(paths, cld.Path)
	}
	return paths
}
