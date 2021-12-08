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
	Short:   "compile a capsule into a deployable form",
	Aliases: []string{"cc"},
	Args:    cobra.MaximumNArgs(1),
	Run:     compile,
}
var srcPath string

func compile(_ *cobra.Command, args []string) {
	if len(args) > 0 {
		srcPath = args[0]
	}
	_, err := capsule.Parse(srcPath, capsule.DefaultConfig())
	if err != nil {
		logrus.Fatal(err)
	}
}
