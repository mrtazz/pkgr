package manifest

import (
	"crypto/sha256"
	"encoding/hex"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Manifest represents a +MANIFEST file
type Manifest struct {
	Arch       string
	Name       string
	Version    string
	Comment    string
	Desc       string
	Origin     string
	Maintainer string
	WWW        string
	Prefix     string
	Files      map[string]string
	Scripts    map[string]string
}

// FromFile reads a manifest from a file on disk
func FromFile(path string) (*Manifest, error) {
	ret := &Manifest{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ret, err
	}

	err = yaml.Unmarshal(data, ret)

	return ret, err
}

// Write writes a manifest to a file
func (m *Manifest) Write(path string) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(data), 0640)
}

// AddFilesFromDir adds files from a directory to the manifest data
func (m *Manifest) AddFilesFromDir(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			chksum, err := sha256sum(path)
			if err != nil {
				log.Printf("Unable to get checksum for %s: %s", path, err.Error())
				return nil
			}
			m.Files[path] = chksum
		}
		return nil
	})
	return err
}

func sha256sum(file string) (string, error) {
	hasher := sha256.New()
	s, err := ioutil.ReadFile(file)
	hasher.Write(s)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
