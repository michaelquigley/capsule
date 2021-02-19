package main

import (
	"bytes"
	"fmt"
	"github.com/michaelquigley/pfxlog"
	"github.com/sirupsen/logrus"
	"github.com/yuin/goldmark"
)

func init() {
	pfxlog.Global(logrus.InfoLevel)
	pfxlog.SetPrefix("github.com/michaelquigley")
}

func main() {
	var buf bytes.Buffer
	source := "# Oh, Wow!"
	if err := goldmark.Convert([]byte(source), &buf); err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes()))
}
