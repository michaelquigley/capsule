package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"sort"
)

type StructureDef struct {
	Models []*StructureModel
}

type StructureModel struct {
	Id      string
	Builder StructureBuilder
	Def     interface{}
}

func LoadStructureDef(path string) (*StructureDef, error) {
	def := &StructureDef{}
	if err := cf.BindYaml(def, path, CfOptions()); err != nil {
		return nil, errors.Wrapf(err, "error loading structure def from '%v' (%v)", path, err)
	}
	return def, nil
}

type StructureBuilder interface {
	Build(rootPath string, node *Node, prev interface{}) (interface{}, error)
}

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
	}
	return paths
}

func init() {
	CfOptions().AddFlexibleSetter("timeline", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &TimelineStructureBuilder{}, nil
	})
}
