package main

import (
	"os"
)

//go:generate go-assets-builder -p server -s="/assets/" -o ../../server/bindata.go ../../assets

func main() {
	os.Exit(Run(os.Args[1:]))
}
