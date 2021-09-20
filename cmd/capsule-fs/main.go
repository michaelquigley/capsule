package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
	"io/fs"
	"os"
)

func init() {
	pfxlog.GlobalInit(logrus.InfoLevel, pfxlog.DefaultOptions().SetTrimPrefix("github.com/michaelquigley/"))
}

func main() {
	dirFs := os.DirFS(".")
	err := fs.WalkDir(dirFs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fi, err := d.Info()
		if err != nil {
			return err
		}
		logrus.Infof("[%v] = %v, %v", path, d.Type(), fi.Size())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
