package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/csaw-ai/csaw/internal/mount"
	"github.com/csaw-ai/csaw/internal/profiles"
	"github.com/csaw-ai/csaw/internal/runtime"
	"github.com/csaw-ai/csaw/internal/sources"
	"github.com/csaw-ai/csaw/internal/workspace"
)

type promptConflictResolver struct {
	cmd      *cobra.Command
	forceAll bool
	skipAll  bool
}

func (r promptConflictResolver) Resolve(conflict mount.Conflict) (mount.ConflictAction, error) {
	if r.forceAll {
		return mount.ConflictOverwrite, nil
	}
	if r.skipAll {
		return mount.ConflictSkip, nil
	}

	info, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
	if info.Mode()&os.ModeCharDevice == 0 {
		return "", errors.New("conflict requires interaction; rerun with --force or --skip-conflicts")
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprintf(r.cmd.OutOrStdout(), "%s already exists. [o]verwrite [s]kip [d]iff > ", conflict.RelativePath)
		answer, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		answer = strings.TrimSpace(strings.ToLower(answer))
		switch answer {
		case "o", "overwrite":
			return mount.ConflictOverwrite, nil
		case "s", "skip":
			return mount.ConflictSkip, nil
		case "d", "diff":
			diffCmd := exec.Command("git", "diff", "--no-index", "--", conflict.TargetPath, conflict.SourcePath)
			diffCmd.Stdout = r.cmd.OutOrStdout()
			diffCmd.Stderr = r.cmd.ErrOrStderr()
			if err := diffCmd.Run(); err != nil {
				var exitErr *exec.ExitError
				if !errors.As(err, &exitErr) || exitErr.ExitCode() != 1 {
					return "", err
				}
			}
		default:
			fmt.Fprintln(r.cmd.OutOrStdout(), "answer with overwrite, skip, or diff")
		}
	}
}

func collectMountEntries(manager sources.Manager, paths runtime.Paths, selection mount.Selection) ([]mount.SourceEntry, error) {
	catalog, err := manager.ExistingCatalog()
	if err != nil {
		return nil, err
	}

	if selection.Profile != "" {
		resolver, err := profiles.NewCatalogResolver(paths, catalog)
		if err != nil {
			return nil, err
		}
		profile, err := resolver.Resolve(selection.Profile)
		if err != nil {
			return nil, err
		}
		selection.IncludePatterns = append(append([]string(nil), profile.Include...), selection.IncludePatterns...)
		selection.ExcludePatterns = append(append([]string(nil), profile.Exclude...), selection.ExcludePatterns...)
		selection.IncludeIgnored = selection.IncludeIgnored || profile.IncludeIgnored
	}

	var entries []mount.SourceEntry
	for _, source := range catalog {
		sourceEntries, err := mount.EnumerateSourceEntries(source)
		if err != nil {
			return nil, err
		}
		if !selection.IncludeIgnored {
			patterns, err := mount.ReadIgnorePatterns(source.Root)
			if err != nil {
				return nil, err
			}
			sourceEntries, err = mount.ApplyIgnore(sourceEntries, patterns)
			if err != nil {
				return nil, err
			}
		}
		entries = append(entries, sourceEntries...)
	}

	return mount.FilterSourceEntries(entries, selection)
}

func findNamedSource(manager sources.Manager, name string) (sources.Source, error) {
	if name == "personal" {
		return sources.Source{Name: "personal", Kind: sources.KindLocal, Path: manager.Paths.Personal}, nil
	}
	return manager.Get(name)
}

func targetProjectRoot() (string, error) {
	projectRoot, err := runtime.FindRepoRoot(".")
	if err == nil {
		return projectRoot, nil
	}
	return os.Getwd()
}

func entriesFromRestoreState(paths runtime.Paths, projectRoot string) ([]mount.SourceEntry, error) {
	state, err := workspace.ReadRestoreState(paths, projectRoot)
	if err != nil {
		return nil, err
	}
	return mount.EntriesFromMountedState(state), nil
}
