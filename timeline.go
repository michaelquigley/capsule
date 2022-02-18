package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/sirupsen/logrus"
	"sort"
)

func init() {
	CfOptions().AddFlexibleSetter("timeline", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &TimelineStructureBuilder{}, nil
	})
}

type TimelineStructure struct {
	Nodes []*Node
}

type TimelineStructureBuilder struct{}

func (self *TimelineStructureBuilder) Build(_ string, node *Node, prev interface{}) (interface{}, error) {
	var timeline *TimelineStructure
	if prev != nil {
		if v, ok := prev.(*TimelineStructure); ok {
			timeline = v
		}
	}
	if timeline == nil {
		timeline = &TimelineStructure{}
	}

	timeline.Nodes = append(timeline.Nodes, self.inventory(node)...)
	sort.Slice(timeline.Nodes, func(i, j int) bool {
		return timeline.Nodes[i].Path < timeline.Nodes[j].Path
	})

	return timeline, nil
}

func (self *TimelineStructureBuilder) inventory(node *Node) []*Node {
	var nodes []*Node
	for _, cld := range node.Children {
		nodes = append(nodes, cld)
		logrus.Debugf("added node '%v' to timeline", cld.Path)
	}
	return nodes
}
