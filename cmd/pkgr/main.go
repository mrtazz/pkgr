package main

import (
	"flag"
	"github.com/mrtazz/pkgr/archiver"
	"github.com/mrtazz/pkgr/manifest"
	"log"
)

type options struct {
	Path string
}

func main() {
	opts := parseOpts()
	m, err := manifest.FromFile(opts.Path + "/+MANIFEST")
	archiver.Archive(opts.Path)
}

func parseOpts() *options {
	ret := &options{}
	flag.StringVar(&ret.Path, "path", "", "directory to build package from")

	flag.Parse()
	return ret
}
