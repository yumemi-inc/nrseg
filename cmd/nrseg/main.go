package main

import (
	"os"

	"github.com/yumemi-inc/nrseg"
)

var (
	Version  = "devel"
	Revision = "unset"
)

func main() {
	if err := nrseg.Run(os.Args, os.Stdout, os.Stderr, Version, Revision); err != nil {
		os.Exit(1)
	}
}
