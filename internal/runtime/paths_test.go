package runtime

import (
	"path/filepath"
	"testing"
)

func TestNormalizeRegistryPath(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: `./skills/debugging`, want: "skills/debugging"},
		{input: `skills\\debugging\\`, want: "skills/debugging"},
		{input: `skills//debugging///notes.md`, want: "skills/debugging/notes.md"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			if got := NormalizeRegistryPath(test.input); got != test.want {
				t.Fatalf("NormalizeRegistryPath(%q) = %q, want %q", test.input, got, test.want)
			}
		})
	}
}

func TestBuildPaths(t *testing.T) {
	root := filepath.Join(string(filepath.Separator), "tmp", "example")
	paths := BuildPaths(root)

	if paths.Root != root {
		t.Fatalf("root = %q, want %q", paths.Root, root)
	}
	if paths.Sources != filepath.Join(root, SourcesDirName) {
		t.Fatalf("sources path mismatch: %q", paths.Sources)
	}
	if paths.Config != filepath.Join(root, ConfigFileName) {
		t.Fatalf("config path mismatch: %q", paths.Config)
	}
	if paths.State != filepath.Join(root, StateDirName) {
		t.Fatalf("state path mismatch: %q", paths.State)
	}
}
