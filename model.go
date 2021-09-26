package capsule

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
	"path/filepath"
)

type Model struct {
	SrcPath string
	Root    *Node
}

type Node struct {
	Path       string
	Capsule    *Capsule
	Properties []*Property
	Parent     *Node
	Children   []*Node
}

type Capsule struct {
	Version   string
	Structure string
}

type Property struct {
	Name string
	Type string
}

// Parse a source path into a Model.
//
func Parse(srcPath string) (*Model, error) {
	srcFs := os.DirFS(srcPath)

	// Inventory sources
	pv := &parseVisitor{make(map[string]*Node)}
	if err := fs.WalkDir(srcFs, ".", pv.visit); err != nil {
		return nil, errors.Wrap(err, "parse error")
	}

	return nil, nil
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
