package main

import (
	"testing"

	"github.com/blang/semver"
)

// TODO: Properly test all 4 different git repo url formats (https://, git://, git+ssh://, user/repo for GitHub)

func TestVersion(t *testing.T) {
	_, err := semver.Make(Version)
	if err != nil {
		t.Errorf("Version failed to validate with error: %s", err)
	}
}
