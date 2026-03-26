package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestDiffCommand(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git not available")
	}
	if canSymlink, reason := canCreateDiffSymlink(); !canSymlink {
		t.Skip(reason)
	}

	root := t.TempDir()
	source := filepath.Join(root, "source.txt")
	link := filepath.Join(root, "mounted.txt")
	if err := os.WriteFile(source, []byte("source\n"), 0o644); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}
	if err := os.Symlink(source, link); err != nil {
		t.Fatalf("Symlink() error = %v", err)
	}

	cmd := newDiffCommand()
	cmd.SetArgs([]string{link})
	if err := cmd.Execute(); err != nil {
		t.Fatalf("Execute() error = %v", err)
	}
}

func canCreateDiffSymlink() (bool, string) {
	root, err := os.MkdirTemp("", "csaw-diff-*")
	if err != nil {
		return false, err.Error()
	}
	defer os.RemoveAll(root)

	target := filepath.Join(root, "target")
	link := filepath.Join(root, "link")
	if err := os.WriteFile(target, []byte("x"), 0o644); err != nil {
		return false, err.Error()
	}
	if err := os.Symlink(target, link); err != nil {
		return false, err.Error()
	}
	return true, ""
}
