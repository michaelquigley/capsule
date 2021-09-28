package capsule

import (
	"github.com/karrick/godirwalk"
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
func Parse(srcPath string, cfg *Config) (model *Model, err error) {
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
	pv := &parseVisitor{cfg, srcPath, make(map[string]*Node)}
	if err := godirwalk.Walk(srcPath, &godirwalk.Options{Callback: pv.visit, Unsorted: true}); err != nil {
		return nil, errors.Wrap(err, "parse error")
	}

	// Link
	model.Root = linkNodes(pv.index)

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
	cfg     *Config
	srcPath string
	index   map[string]*Node
}

func (pv *parseVisitor) visit(path string, de *godirwalk.Dirent) error {
	logrus.Infof("visiting '%v'", path)

	dir := filepath.Dir(path)

	if de.IsDir() {
		// Empty directories
		node, found := pv.index[dir]
		if !found {
			node = &Node{
				Path: dir,
			}
			pv.index[dir] = node
		}
	} else if de.IsRegular() {
		if filepath.Base(path) != ".capsule" {
			node, found := pv.index[dir]
			if !found {
				node = &Node{
					Path: dir,
				}
				pv.index[dir] = node
			}
			prop := &Property{Name: filepath.Base(path)}
			typeId, typeFound := pv.cfg.PropertyType(path, de)
			if typeFound {
				prop.Type = typeId
			}
			node.Properties = append(node.Properties, prop)
		}
	}

	return nil
}

func linkNodes(index map[string]*Node) *Node {
	logrus.Infof("index = %v", index)

	var root *Node
	for _, node := range index {
		parent, found := index[filepath.Dir(node.Path)]
		if !found {
			root = node
			root.Path = "."
		} else {
			node.Parent = parent
			parent.Children = append(parent.Children, node)
			node.Path = filepath.Base(node.Path)
		}
	}
	return root
}
