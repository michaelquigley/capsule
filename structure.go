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

// StructureBuilder is used to create structures suited for capturing additional details about a capsule structure.
// When a structure definition is encountered in the tree, the models are built using the Build method below. Build
// takes a previous version of the identified model, allows additional transformation to be done to it, and emits an
// altered version to insert back into the model.
//
type StructureBuilder interface {
	Build(rootPath string, node *Node, prev interface{}) (interface{}, error)
}
