package capsule

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
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
	if err := fs.WalkDir(srcFs, ".", parseVisitor); err != nil {
		return nil, errors.Wrap(err, "parse error")
	}
	return nil, nil
}

func parseVisitor(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	if !d.IsDir() {
		logrus.Infof("'%v'", path)
	}
	return nil
}
