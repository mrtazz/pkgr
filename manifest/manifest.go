package manifest

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Manifest represents a +MANIFEST file
type Manifest struct {
	Arch       string            `json:"arch"`
	Name       string            `json:"name"`
	Version    string            `json:"version"`
	Comment    string            `json:"comment"`
	Desc       string            `json:"desc"`
	Origin     string            `json:"origin"`
	Maintainer string            `json:"maintainer"`
	WWW        string            `json:"www"`
	Prefix     string            `json:"prefix"`
	Files      map[string]string `json:"files,omitempty"`
	Scripts    map[string]string `json:"scripts,omitempty"`
}

// FromFile reads a manifest from a file on disk
func FromFile(path string) (*Manifest, error) {
	ret := &Manifest{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return ret, err
	}

	err = yaml.Unmarshal(data, ret)

	ret.Origin = fmt.Sprintf("pkgr/%s", ret.Name)
	ret.Arch = fmt.Sprintf("%s:%s:%s", runtime.GOOS, "11", runtime.GOARCH)
	ret.Prefix = "/"

	return ret, err
}

// Write writes a manifest to a file
func (m *Manifest) Write(path string) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	data, err = yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(data), 0640)
}

func (m *Manifest) WriteCompact(path string) error {
	data, err := yaml.Marshal(m)
	if err != nil {
		return err
	}
	data, err = yaml.YAMLToJSON(data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, []byte(data), 0640)
}

// AddFilesFromDir adds files from a directory to the manifest data
func (m *Manifest) AddFilesFromDir(path string) error {
	m.Files = make(map[string]string)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, "MANIFEST") {
			return nil
		}
		if !info.IsDir() {
			chksum, err := sha256sum(path)
			if err != nil {
				log.Printf("Unable to get checksum for %s: %s", path, err.Error())
				return nil
			}
			m.Files[fmt.Sprintf("/%s", path)] = chksum
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
