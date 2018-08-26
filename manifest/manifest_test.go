package manifest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSha256Sum(t *testing.T) {
	shasum, err := sha256sum("checksum.test")
	assert.Equal(t, nil, err, "shasum should have worked")
	assert.Equal(t, "157925f6a0219aeeffb73f428fa36f79a8ba91bb226f623ff861778432f7810f", shasum, "shasum should be correct")

	// non existent file throws error
	shasum, err = sha256sum("checksum.nothere")
	assert.Equal(t, "", shasum, "shasum should be empty string")
	assert.Equal(t, "open checksum.nothere: no such file or directory", err.Error(), "shasum should be correct")
}

func TestFreeBSDMajorVersion(t *testing.T) {
	FreeBSDVersionCommand = "./fake-freebsd-version.sh"
	version, err := getFreeBSDMajorVersion()
	assert.Equal(t, nil, err, "version parsing should have worked")
	assert.Equal(t, "11", version, "version should be 11")
}
