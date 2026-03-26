package mount

import "testing"

func TestFilterEntries(t *testing.T) {
	planner := NewPlanner()
	entries := []string{
		"agents/base.md",
		"agents/go.md",
		"skills/debugging/SKILL.md",
		"skills/experimental/SKILL.md",
	}

	selection := Selection{
		IncludePatterns: []string{"agents", "skills/**"},
		ExcludePatterns: []string{"skills/experimental/**"},
	}

	filtered, err := planner.Filter(entries, selection)
	if err != nil {
		t.Fatalf("Filter() error = %v", err)
	}

	if got, want := len(filtered), 3; got != want {
		t.Fatalf("len(filtered) = %d, want %d (%v)", got, want, filtered)
	}
}

func TestFilterDefaultsToAllWhenNoIncludes(t *testing.T) {
	planner := NewPlanner()
	entries := []string{"AGENTS.md", "skills/debugging/SKILL.md"}

	filtered, err := planner.Filter(entries, Selection{})
	if err != nil {
		t.Fatalf("Filter() error = %v", err)
	}

	if got, want := len(filtered), len(entries); got != want {
		t.Fatalf("len(filtered) = %d, want %d", got, want)
	}
}
