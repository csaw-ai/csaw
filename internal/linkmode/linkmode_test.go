package linkmode

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateAndVerifySymlink(t *testing.T) {
	if Detect() != Symlink {
		t.Skip("symlinks not available")
	}

	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	target := filepath.Join(dir, "target.md")
	if err := os.WriteFile(source, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Create(Symlink, source, target); err != nil {
		t.Fatal(err)
	}

	if !IsLink(Symlink, target, source) {
		t.Fatal("IsLink should return true for symlink")
	}

	equal := func(a, b string) bool { return a == b }
	healthy, resolved := Verify(Symlink, target, source, equal)
	if !healthy {
		t.Fatalf("Verify returned unhealthy, resolved=%q", resolved)
	}

	got, err := ReadTarget(Symlink, target)
	if err != nil {
		t.Fatal(err)
	}
	if got != source {
		t.Fatalf("ReadTarget = %q, want %q", got, source)
	}
}

func TestCreateAndVerifyHardlink(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	target := filepath.Join(dir, "target.md")
	if err := os.WriteFile(source, []byte("hello"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Create(Hardlink, source, target); err != nil {
		t.Fatal(err)
	}

	if !IsLink(Hardlink, target, source) {
		t.Fatal("IsLink should return true for hardlink to same file")
	}

	equal := func(a, b string) bool { return a == b }
	healthy, _ := Verify(Hardlink, target, source, equal)
	if !healthy {
		t.Fatal("Verify should return healthy for valid hardlink")
	}

	// Verify that edits propagate (live-update property)
	if err := os.WriteFile(source, []byte("updated"), 0o644); err != nil {
		t.Fatal(err)
	}
	content, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	if string(content) != "updated" {
		t.Fatalf("hardlink content = %q, want %q", string(content), "updated")
	}
}

func TestVerifyDetectsDriftedHardlink(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	target := filepath.Join(dir, "target.md")
	if err := os.WriteFile(source, []byte("original"), 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Create(Hardlink, source, target); err != nil {
		t.Fatal(err)
	}

	// Simulate git pull replacing the source file (delete + recreate breaks hardlink)
	if err := os.Remove(source); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(source, []byte("new version"), 0o644); err != nil {
		t.Fatal(err)
	}

	equal := func(a, b string) bool { return a == b }
	healthy, _ := Verify(Hardlink, target, source, equal)
	if healthy {
		t.Fatal("Verify should detect drifted hardlink after source replacement")
	}
}

func TestIsLinkReturnsFalseForRegularFile(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "source.md")
	target := filepath.Join(dir, "target.md")
	if err := os.WriteFile(source, []byte("a"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(target, []byte("b"), 0o644); err != nil {
		t.Fatal(err)
	}

	if IsLink(Symlink, target, source) {
		t.Fatal("regular file should not be detected as symlink")
	}
	if IsLink(Hardlink, target, source) {
		t.Fatal("unrelated regular file should not be detected as hardlink")
	}
}
