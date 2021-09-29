package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
)

// StructuralDirective is a strategy that is used to build structure from a Node tree.
//
type StructuralDirective interface {
	Build(*Node, map[string]interface{}) error
}

func LoadStructure(path string) (*Structure, error) {
	opt := cf.DefaultOptions()
	opt.AddFlexibleSetter("timeline", func(v interface{}, opt *cf.Options) (interface{}, error) {
		return &TimelineStructure{}, nil
	})
	s := &Structure{}
	if err := cf.BindYaml(s, path, opt); err != nil {
		return nil, errors.Wrapf(err, "error loading structure from '%v'", path)
	}
	return s, nil
}

// Structure model
//
type Structure struct {
	Models []interface{}
}

type TimelineStructure struct{}

func (ts *TimelineStructure) Build(n *Node, structure map[string]interface{}) error {
	return nil
}
