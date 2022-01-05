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

func (self *TimelineStructureBuilder) Build(rootPath string, node *Node, prev interface{}) (interface{}, error) {
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

func (self *TimelineStructureBuilder) inventory(node *Node) []string {
	var paths []string
	for _, cld := range node.Children {
		paths = append(paths, cld.Path)
		logrus.Debugf("added node '%v' to timeline", cld.Path)
	}
	return paths
}

func init() {
	CfOptions().AddFlexibleSetter("timeline", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &TimelineStructureBuilder{}, nil
	})
}
