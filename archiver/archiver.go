package archiver

import (
	"archive/tar"
	"fmt"
	"github.com/ulikunitz/xz"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func addFile(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	tarpath := path
	if !strings.Contains(path, "MANIFEST") {
		tarpath = fmt.Sprintf("/%s", path)
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = tarpath
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

// Archive creates a gzip compressed tar archive of the given directory
func Archive(tarball, dir string) error {
	// set up the output file
	file, err := os.Create(tarball)
	if err != nil {
		return err
	}
	defer file.Close()
	// set up the xz writer
	xzw, err := xz.NewWriter(file)
	if err != nil {
		return err
	}
	defer xzw.Close()
	tw := tar.NewWriter(xzw)
	defer tw.Close()
	// grab the paths that need to be added in
	paths := make([]string, 0, 10)
	paths = append(paths, "+MANIFEST")
	paths = append(paths, "+COMPACT_MANIFEST")
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	// add each file as needed into the current tar archive
	for i := range paths {
		if err := addFile(tw, paths[i]); err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}
