package docs

import (
	"path/filepath"
	"testing"
)

func repoRoot(t *testing.T) string {
	t.Helper()

	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatalf("filepath.Abs() error = %v", err)
	}
	return root
}

func TestRepositoryScaffold(t *testing.T) {
	root := repoRoot(t)

	if err := ValidateAgentsLinks(root); err != nil {
		t.Fatalf("ValidateAgentsLinks() error = %v", err)
	}
	if err := ValidateActiveExecPlans(root); err != nil {
		t.Fatalf("ValidateActiveExecPlans() error = %v", err)
	}
	if err := ValidateSkills(root); err != nil {
		t.Fatalf("ValidateSkills() error = %v", err)
	}
	if err := ValidatePublicRepoContent(root); err != nil {
		t.Fatalf("ValidatePublicRepoContent() error = %v", err)
	}
}
