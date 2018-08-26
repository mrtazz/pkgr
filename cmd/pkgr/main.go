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
	Manifest     string
	PrintVersion bool
}

const (
	MANIFEST         = "+MANIFEST"
	COMPACT_MANIFEST = "+COMPACT_MANIFEST"
)

func main() {
	opts := parseOpts()
	if opts.PrintVersion {
		fmt.Printf("version %s, built by %s with %s ", version, builder, goversion)
		os.Exit(0)
	}
	err := validateOpts(opts)
	if err != nil {
		fmt.Println(err.Error())
		flag.Usage()
		os.Exit(1)
	}
	m, err := manifest.FromFile(opts.Manifest)
	if err != nil {
		log.Printf("Unable to load manifest: %s", err.Error())
		os.Exit(1)
	}
	err = m.Write(COMPACT_MANIFEST)
	if err != nil {
		log.Printf("Unable to write compact manifest: %s", err.Error())
		os.Exit(1)
	}
	err = m.AddFilesFromDir(opts.Path)
	if err != nil {
		log.Printf("Unable to load files from dir: %s", err.Error())
		os.Exit(1)
	}
	err = m.Write(MANIFEST)
	if err != nil {
		log.Printf("Unable to write manifest: %s", err.Error())
		os.Exit(1)
	}
	err = archiver.Archive(fmt.Sprintf("%s-%s.txz", m.Name, m.Version), opts.Path)
	if err != nil {
		fmt.Printf("Unable to write package archive: %s", err.Error())
		os.Exit(1)
	}

	// clean up generated files
	err = os.Remove(COMPACT_MANIFEST)
	if err != nil {
		fmt.Printf("Unable to remove compact manifest: %s", err.Error())
		os.Exit(1)
	}
	err = os.Remove(MANIFEST)
	if err != nil {
		fmt.Printf("Unable to remove manifest: %s", err.Error())
		os.Exit(1)
	}
}

func parseOpts() *options {
	ret := &options{}
	flag.StringVar(&ret.Path, "path", "", "directory to build package from")
	flag.StringVar(&ret.Manifest, "manifest", "", "manifest to read in for package information")
	flag.BoolVar(&ret.PrintVersion, "version", false, "print version and exit")

	flag.Parse()
	return ret
}

func validateOpts(opts *options) error {

	if opts.Path == "" {
		return fmt.Errorf("Please provide a directory to package.")
	}

	if opts.Manifest == "" {
		return fmt.Errorf("Please provide a manifest.")
	}

	return nil
}
