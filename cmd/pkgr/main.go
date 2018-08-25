package main

import (
	"flag"
	"fmt"
	"github.com/mrtazz/pkgr/archiver"
	"github.com/mrtazz/pkgr/manifest"
	"log"
	"os"
)

var (
	version   = ""
	builder   = ""
	goversion = ""
)

type options struct {
	Path         string
	PrintVersion bool
}

func main() {
	opts := parseOpts()
	if opts.PrintVersion {
		fmt.Printf("version %s, built by %s with %s ", version, builder, goversion)
		os.Exit(0)
	}
	m, err := manifest.FromFile(opts.Path + "/+MANIFEST")
	archiver.Archive(opts.Path)
}

func parseOpts() *options {
	ret := &options{}
	flag.StringVar(&ret.Path, "path", "", "directory to build package from")
	flag.BoolVar(&ret.PrintVersion, "version", false, "print version and exit")

	flag.Parse()
	return ret
}
