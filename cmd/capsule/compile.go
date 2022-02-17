package main

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/compiler/static"
	_ "github.com/michaelquigley/capsule/compiler/static"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(newCompileCommand().cmd)
}

type compileCommand struct {
	cmd          *cobra.Command
	capsulePath  string
	resourcePath string
	buildPath    string
}

func newCompileCommand() *compileCommand {
	cc := &compileCommand{
		cmd: &cobra.Command{
			Use:     "compile",
			Short:   "compile a capsule into deployable build",
			Aliases: []string{"cc"},
			Args:    cobra.ExactArgs(0),
		},
	}
	cc.cmd.Run = cc.run
	cc.cmd.Flags().StringVarP(&cc.capsulePath, "capsule", "c", "src/", "capsule path")
	cc.cmd.Flags().StringVarP(&cc.resourcePath, "resource", "r", "resources/", "resource path")
	cc.cmd.Flags().StringVarP(&cc.buildPath, "build", "b", "build/", "build path")
	return cc
}

func (cc *compileCommand) run(_ *cobra.Command, _ []string) {
	m, err := capsule.Parse(cc.capsulePath, capsule.DefaultConfig())
	if err != nil {
		panic(err)
	}

	fmt.Println(m.Dump())

	st := static.New(&static.Config{BuildPath: cc.buildPath, ResourcePath: cc.resourcePath})
	if err := st.Compile(m); err != nil {
		panic(err)
	}
}
