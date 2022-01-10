package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"sort"
)

type TimelineStructure struct {
	Nodes []*Node
}

type TimelineStructureBuilder struct{}

func (self *TimelineStructureBuilder) Build(_ string, node *Node, prev interface{}) (interface{}, error) {
	var prevarr []*Node
	if prev != nil {
		var ok bool
		prevarr, ok = prev.([]*Node)
		if !ok {
			return nil, errors.Errorf("invalid previous state")
		}
	}
	nextarr := make([]*Node, len(prevarr))
	copy(nextarr, prevarr)

	nodearr := self.inventory(node)
	nextarr = append(nextarr, nodearr...)
	sort.Slice(nextarr, func(i, j int) bool {
		return nextarr[i].Path < nextarr[j].Path
	})
	return nextarr, nil
}

func (self *TimelineStructureBuilder) inventory(node *Node) []*Node {
	var nodes []*Node
	for _, cld := range node.Children {
		nodes = append(nodes, cld)
		logrus.Debugf("added node '%v' to timeline", cld.Path)
	}
	return nodes
}

func init() {
	CfOptions().AddFlexibleSetter("timeline", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &TimelineStructureBuilder{}, nil
	})
}
