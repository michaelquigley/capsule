package main

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	dumpSrcPath = "."
	rootCmd.AddCommand(dumpCmd)
}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump a Capsule Model",
	Args:  cobra.MaximumNArgs(1),
	Run:   dump,
}
var dumpSrcPath string

func dump(_ *cobra.Command, args []string) {
	if len(args) > 0 {
		srcPath = args[0]
	}
	model, err := capsule.Parse(srcPath)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Println(capsule.Dump(model))
}
