package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
)

type StructureDef struct {
	Models []*StructureModel
}

type StructureModel struct {
	Id      string
	Builder interface{}
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
