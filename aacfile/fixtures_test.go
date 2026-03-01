package aacfile

import (
	"os"
	"path/filepath"
	"testing"
)

const largeFixtureEnv = "CLIMP_AAC_LARGE_FIXTURE"

func fixturePath(t testing.TB, name string) string {
	t.Helper()

	if name == "BeyondGoodEvil_librivox.m4b" {
		path := os.Getenv(largeFixtureEnv)
		if path == "" {
			t.Skipf("%s not set", largeFixtureEnv)
		}
		return path
	}

	return filepath.Join("..", "testdata", "fixtures", name)
}
