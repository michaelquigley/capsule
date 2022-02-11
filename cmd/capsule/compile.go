package main

import (
	"fmt"
	"github.com/michaelquigley/capsule"
	"github.com/michaelquigley/capsule/compiler"
	_ "github.com/michaelquigley/capsule/compiler/static"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	compileCmd.Flags().StringVarP(&capsulePath, "capsule", "c", ".", "capsule path")
	compileCmd.Flags().StringVarP(&compilerId, "compiler", "C", "static", "compiler id")
	compileCmd.Flags().StringVarP(&compilerVars, "compiler-vars", "V", "", "compiler variables <var>=<val>[,<var>=<val>]")
	rootCmd.AddCommand(compileCmd)
}

var compileCmd = &cobra.Command{
	Use:     "compile",
	Short:   "compile a capsule into a deployable form",
	Aliases: []string{"cc"},
	Args:    cobra.ExactArgs(0),
	Run:     compile,
}
var capsulePath string
var compilerId string
var compilerVars string

func compile(_ *cobra.Command, _ []string) {
	mdl, err := capsule.Parse(capsulePath, capsule.DefaultConfig())
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Println(mdl.Dump())

	if ccFactory, err := compiler.Get(compilerId); err == nil {
		if ccCfg, err := parseCompilerCfg(); err == nil {
			if cc, err := ccFactory.New(ccCfg); err == nil {
				if err := cc.Compile(mdl); err != nil {
					logrus.Fatal(err)
				}
			} else {
				logrus.Fatalf("error creating compiler (%v)", err)
			}
		} else {
			logrus.Fatalf("error parsing compiler config (%v)", err)
		}
	} else {
		logrus.Fatalf("error getting compiler factory (%v)", err)
	}
}

func parseCompilerCfg() (compiler.Config, error) {
	cfg := make(compiler.Config)
	if compilerVars != "" {
		variables := strings.Split(compilerVars, ",")
		for _, variable := range variables {
			sides := strings.Split(strings.TrimSpace(variable), "=")
			if len(sides) != 2 {
				return nil, errors.Errorf("invalid variable def '%v'", variable)
			}
			lside := strings.TrimSpace(sides[0])
			rside := strings.TrimSpace(sides[1])
			cfg[lside] = rside
		}
	}
	return cfg, nil
}
