package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
)

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
	if err := fs.WalkDir(os.DirFS(srcPath), ".", pv.visit); err != nil {
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

func (pv *parseVisitor) visit(path string, de fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	dir := filepath.ToSlash(filepath.Join(pv.srcPath, filepath.Dir(path)))
	path = filepath.ToSlash(filepath.Join(pv.srcPath, path))

	node, found := pv.index[dir]
	if !found {
		node = &Node{
			Path: dir,
		}
		pv.index[dir] = node
		logrus.Infof("pv.index[%v] = %v", dir, node)
	}

	if !de.IsDir() {
		if filepath.Base(path) != ".capsule" && filepath.Base(path) != ".structure" {
			ftr := &Feature{Name: filepath.Base(path)}
			typeId, typeFound := pv.cfg.PropertyType(path, de)
			if typeFound {
				ftr.Type = typeId
			}
			node.Features = append(node.Features, ftr)
		}
		if filepath.Base(path) == ".structure" {
			str, err := LoadStructure(path)
			if err != nil {
				return err
			}
			node.Structure = append(node.Structure, str)
		}
	}

	return nil
}

func linkNodes(index map[string]*Node) *Node {
	var root *Node
	for _, node := range index {
		parentPath := filepath.ToSlash(filepath.Dir(node.Path))
		parent, found := index[parentPath]
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
