package main

import (
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
)

func init() {
	o := pfxlog.DefaultOptions().SetTrimPrefix("github.com/michaelquigley/")
	pfxlog.GlobalInit(logrus.InfoLevel, o)
	logrus.SetFormatter(pfxlog.NewFormatter(o)) // Always formatted output
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
