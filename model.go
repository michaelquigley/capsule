package capsule

import (
	"github.com/karrick/godirwalk"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"path/filepath"
)

const capsuleVersion = "v1"

type Model struct {
	SrcPath string
	Capsule *Capsule
	Root    *Node
}

type Node struct {
	Path       string
	Structure  []*Structure
	Properties []*Property
	Parent     *Node
	Children   []*Node
}

// Capsule metadata
//
type Capsule struct {
	Version string
}

// Structure definition
//
type Structure struct {
	Models []interface{}
}

type Property struct {
	Name string
	Type string
}

// Parse a source path into a Model.
//
func Parse(srcPath string) (model *Model, err error) {
	model = &Model{
		SrcPath: srcPath,
	}

	// Load capsule metadata
	model.Capsule, err = loadCapsule(filepath.Join(srcPath, ".capsule"))
	if err != nil {
		return nil, err
	}
	if model.Capsule.Version != capsuleVersion {
		return nil, errors.Errorf("invalid capsule version '%v', expected '%v'", model.Capsule.Version, capsuleVersion)
	}

	// Inventory sources
	pv := &parseVisitor{srcPath, make(map[string]*Node)}
	if err := godirwalk.Walk(srcPath, &godirwalk.Options{Callback: pv.visit, Unsorted: true}); err != nil {
		return nil, errors.Wrap(err, "parse error")
	}

	return model, nil
}

func loadCapsule(capsulePath string) (*Capsule, error) {
	c := &Capsule{}
	if err := cf.BindYaml(c, capsulePath, cf.DefaultOptions()); err != nil {
		return nil, errors.Wrapf(err, "load capsule path '%v'", capsulePath)
	}
	return c, nil
}

type parseVisitor struct {
	srcPath string
	index   map[string]*Node
}

func (pv *parseVisitor) visit(path string, de *godirwalk.Dirent) error {
	dir := filepath.Dir(path)

	if de.IsRegular() {
		node, found := pv.index[dir]
		if !found {
			node = &Node{
				Path: dir,
			}
			pv.index[dir] = node
		}
		node.Properties = append(node.Properties, &Property{Name: filepath.Base(path)})
	}

	return nil
}
