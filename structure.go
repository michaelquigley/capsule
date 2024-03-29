package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
)

type StructureDef struct {
	Models []*ModelDef
}

type ModelDef struct {
	Id      string
	Builder interface{}
}

func LoadStructureDef(path string) (*StructureDef, error) {
	def := &StructureDef{}

	options := cf.DefaultOptions()
	for k, v := range structureBuilderRegistry {
		options.AddFlexibleSetter(k, v)
	}

	if err := cf.BindYaml(def, path, options); err != nil {
		return nil, errors.Wrapf(err, "error loading structure def from '%v' (%v)", path, err)
	}

	return def, nil
}

// Structure defines the interface that must be implemented by all structure models.
//
type Structure interface {
	// Dump emits a string representation of the structure model for use in dump debugging output.
	//
	Dump() string
}

// StructureBuilder is used to create structures suited for capturing additional details about a capsule structure.
// When a structure definition is encountered in the tree, the models are built using the Build method below. Build
// takes a previous version of the identified model, allows additional transformation to be done to it, and emits an
// altered version to insert back into the model.
//
type StructureBuilder interface {
	Build(rootPath string, node *Node, prev Structure) (Structure, error)
}

func RegisterStructureBuilder(id string, fs cf.FlexibleSetter) {
	if structureBuilderRegistry == nil {
		structureBuilderRegistry = make(map[string]cf.FlexibleSetter)
	}
	structureBuilderRegistry[id] = fs
}

var structureBuilderRegistry map[string]cf.FlexibleSetter

const StructureFeature = "structure.yaml"
