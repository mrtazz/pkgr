package archiver

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestArchive(t *testing.T) {
	os.Create("+MANIFEST")
	os.Create("+COMPACT_MANIFEST")
	err := Archive("foo.txz", "fixtures/")
	assert.Equal(t, nil, err, "archiving should have worked")
	os.Remove("+MANIFEST")
	os.Remove("+COMPACT_MANIFEST")
}
