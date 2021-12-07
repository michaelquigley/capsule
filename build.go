package capsule

import (
	"github.com/pkg/errors"
	"reflect"
)

func (model *Model) Build() error {
	if err := model.buildStructureFor(model.Root); err != nil {
		return errors.Wrap(err, "error building structure")
	}
	return nil
}

func (model *Model) buildStructureFor(n *Node) error {
	for _, child := range n.Children {
		if err := model.buildStructureFor(child); err != nil {
			return errors.Wrapf(err, "error building structure for '%v'", child.FullPath())
		}
	}
	for _, structure := range n.Structure {
		for _, sm := range structure.Models {
			if err := sm.(StructuralDirective).Build(n, model.Structure); err != nil {
				return errors.Wrapf(err, "error running structural directive '%v' for '%v'", reflect.TypeOf(sm).Name(), n.FullPath())
			}
		}
	}
	return nil
}
