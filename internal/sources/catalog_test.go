package sources

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/csaw-ai/csaw/internal/runtime"
)

type recordingGit struct {
	calls   [][]string
	outputs map[string]string
}

func (g *recordingGit) Run(ctx context.Context, cwd string, args ...string) (string, error) {
	_ = ctx
	call := append([]string{cwd}, args...)
	g.calls = append(g.calls, call)
	if g.outputs == nil {
		return "", nil
	}
	return g.outputs[joinArgs(args)], nil
}

func joinArgs(values []string) string {
	return filepath.ToSlash(filepath.Join(values...))
}

func TestCatalogIncludesPersonal(t *testing.T) {
	paths := runtime.BuildPaths(filepath.Join(t.TempDir(), ".csaw"))
	manager := Manager{Paths: paths}

	catalog, err := manager.Catalog()
	if err != nil {
		t.Fatalf("Catalog() error = %v", err)
	}
	if len(catalog) == 0 || catalog[0].Name != "personal" {
		t.Fatalf("Catalog() = %#v, want personal source", catalog)
	}
}

func TestPushPersonal(t *testing.T) {
	root := t.TempDir()
	paths := runtime.BuildPaths(filepath.Join(root, ".csaw"))
	personalGit := filepath.Join(paths.Personal, ".git")
	if err := os.MkdirAll(personalGit, 0o755); err != nil {
		t.Fatalf("MkdirAll() error = %v", err)
	}

	git := &recordingGit{
		outputs: map[string]string{
			joinArgs([]string{"status", "--porcelain"}): " M AGENTS.md",
		},
	}
	manager := Manager{Paths: paths, Git: git}

	if err := manager.PushPersonal(context.Background(), "test commit"); err != nil {
		t.Fatalf("PushPersonal() error = %v", err)
	}

	if got, want := len(git.calls), 4; got != want {
		t.Fatalf("len(calls) = %d, want %d", got, want)
	}
}
