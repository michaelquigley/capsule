package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
)

func Parse(srcPath string, cfg *Config) (model *Model, err error) {
	srcPath, err = filepath.Abs(srcPath)
	if err != nil {
		return nil, errors.Wrapf(err, "error making src path '%v' absolute", srcPath)
	}
	srcPath = filepath.ToSlash(srcPath)

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
	}

	if !de.IsDir() {
		ftr := &Feature{
			Name:       filepath.Base(path),
			Attributes: pv.cfg.GetAttributes(path, de),
		}
		if de.Name() == ".structure" {
			if str, err := LoadStructureDef(path); err == nil {
				ftr.Object = str
				logrus.Debugf("structure loaded from '%v' with %d models", path, len(str.Models))
			} else {
				return errors.Wrapf(err, "error loading structure from '%v'", path)
			}
		}
		node.Features = append(node.Features, ftr)
		logrus.Debugf("'%v': %v", filepath.Base(path), ftr.Attributes.String())
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
