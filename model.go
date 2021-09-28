package capsule

import (
	"github.com/michaelquigley/cf"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
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
func Parse(srcPath string) (*Model, error) {
	srcFs := os.DirFS(srcPath)

	// Load capsule metadata
	c, err := loadCapsule(filepath.Join(srcPath, ".capsule"))
	if err != nil {
		return nil, err
	}
	if c.Version != capsuleVersion {
		return nil, errors.Errorf("invalid capsule version '%v', expected '%v'", c.Version, capsuleVersion)
	}

	// Inventory sources
	pv := &parseVisitor{make(map[string]*Node)}
	if err := fs.WalkDir(srcFs, ".", pv.visit); err != nil {
		return nil, errors.Wrap(err, "parse error")
	}

	return nil, nil
}

func loadCapsule(capsulePath string) (*Capsule, error) {
	c := &Capsule{}
	if err := cf.BindYaml(c, capsulePath, cf.DefaultOptions()); err != nil {
		return nil, errors.Wrapf(err, "load capsule path '%v'", capsulePath)
	}
	return c, nil
}

type parseVisitor struct {
	index map[string]*Node
}

func (pv *parseVisitor) visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		dir, err := filepath.Abs(filepath.Dir(path))
		if err != nil {
			return errors.Wrap(err, "absolute")
		}
		logrus.Infof("%v:(%v)", dir, filepath.Base(path))
	}
	return nil
}
