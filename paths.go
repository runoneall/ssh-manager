package main

import (
	"os"
	"path"
)

func getProgramRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

var PROGRAM_ROOT string = getProgramRoot()
var CONFIG_FOLDER = path.Join(PROGRAM_ROOT, "config")
