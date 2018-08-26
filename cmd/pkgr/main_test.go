package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateOptions(t *testing.T) {

	opts := &options{Path: "/bla", Manifest: "MANIFEST"}
	err := validateOpts(opts)
	assert.Equal(t, nil, err, "options should be valid")

	opts = &options{Path: "/bla"}
	err = validateOpts(opts)
	assert.Equal(t, "Please provide a manifest.", err.Error(), "should need a Manifest")

	opts = &options{Manifest: "MANIFEST"}
	err = validateOpts(opts)
	assert.Equal(t, "Please provide a directory to package.", err.Error(), "should need a path")
}
