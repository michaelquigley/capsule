package main

import (
	"github.com/michaelquigley/capsule"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	srcPath = "."
	rootCmd.AddCommand(compileCmd)
}

var compileCmd = &cobra.Command{
	Use:     "compile",
	Short:   "Compile a Capsule",
	Aliases: []string{"cc"},
	Args:    cobra.MaximumNArgs(1),
	Run:     compile,
}
var srcPath string

func compile(_ *cobra.Command, args []string) {
	if len(args) > 0 {
		srcPath = args[0]
	}
	model, err := capsule.Parse(srcPath, capsule.DefaultConfig())
	if err != nil {
		logrus.Fatal(err)
	}
	if err := model.Build(); err != nil {
		logrus.Fatal(err)
	}
}
