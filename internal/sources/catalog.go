package sources

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/csaw-ai/csaw/internal/runtime"
)

var ErrNothingToPush = errors.New("nothing to push")

type CatalogSource struct {
	Name string
	Kind string
	Root string
}

func PersonalSource(paths runtime.Paths) CatalogSource {
	return CatalogSource{
		Name: "personal",
		Kind: KindLocal,
		Root: paths.Personal,
	}
}

func (m Manager) Catalog() ([]CatalogSource, error) {
	cfg, err := m.Load()
	if err != nil {
		return nil, err
	}

	catalog := []CatalogSource{PersonalSource(m.Paths)}
	for _, source := range cfg.Sources {
		catalog = append(catalog, CatalogSource{
			Name: source.Name,
			Kind: source.Kind,
			Root: source.CheckoutPath(m.Paths),
		})
	}

	sort.Slice(catalog, func(i, j int) bool { return catalog[i].Name < catalog[j].Name })
	return catalog, nil
}

func (m Manager) ExistingCatalog() ([]CatalogSource, error) {
	catalog, err := m.Catalog()
	if err != nil {
		return nil, err
	}

	filtered := catalog[:0]
	for _, source := range catalog {
		info, err := os.Stat(source.Root)
		if err != nil || !info.IsDir() {
			continue
		}
		filtered = append(filtered, source)
	}

	return filtered, nil
}

func (m Manager) PushPersonal(ctx context.Context, message string) error {
	root := m.Paths.Personal
	if message == "" {
		message = "csaw: update personal registry"
	}

	gitDir := filepath.Join(root, ".git")
	if _, err := os.Stat(gitDir); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("personal registry is not a git repository: %s", root)
		}
		return err
	}

	status, err := m.Git.Run(ctx, root, "status", "--porcelain")
	if err != nil {
		return err
	}
	if strings.TrimSpace(status) == "" {
		return ErrNothingToPush
	}

	if _, err := m.Git.Run(ctx, root, "add", "-A"); err != nil {
		return err
	}
	if _, err := m.Git.Run(ctx, root, "commit", "-m", message); err != nil {
		return err
	}
	_, err = m.Git.Run(ctx, root, "push")
	return err
}
