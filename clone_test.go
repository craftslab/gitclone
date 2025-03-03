//go:build linux

package main

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	configTest = Config{
		Clone: Clone{
			Depth:        cloneDepth,
			SingleBranch: true,
		},
	}
)

func TestCloneRepo(t *testing.T) {
	ctx := context.Background()

	repoUrl = "https://android.googlesource.com/platform/build/soong"
	destDir, _ = os.MkdirTemp("", "clone-repo-test")

	err := cloneRepo(ctx, &configTest)
	assert.Equal(t, nil, err)
}
