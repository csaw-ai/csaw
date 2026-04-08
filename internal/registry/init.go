package registry

import (
	"context"
	"os"
	"path/filepath"

	"github.com/csaw-ai/csaw/internal/git"
)

type InitResult struct {
	Path string
	Name string
}

var starterProfile = `default:
  description: Mount everything
  include:
    - agents/**
    - skills/**
`

var starterIgnore = `# Patterns listed here are excluded from mounting by default.
# Use --include-ignored to override.
`

var starterAgent = `# Agent Instructions

Add your base coding preferences, conventions, and rules here.
This file will be mounted as AGENTS.md in your projects.
`

var starterSkill = `---
name: commit-message
description: Write clear, conventional commit messages
---

When writing commit messages:

- Use the imperative mood ("Add feature" not "Added feature")
- Keep the subject line under 72 characters
- Separate subject from body with a blank line
- Use the body to explain what and why, not how
- Reference issues and PRs where relevant
`

func Init(ctx context.Context, g git.Git, dir string, name string) (InitResult, error) {
	absDir, err := filepath.Abs(dir)
	if err != nil {
		return InitResult{}, err
	}

	if name == "" {
		name = filepath.Base(absDir)
	}

	if err := os.MkdirAll(absDir, 0o755); err != nil {
		return InitResult{}, err
	}

	for _, sub := range []string{"agents", "skills/commit-message"} {
		if err := os.MkdirAll(filepath.Join(absDir, sub), 0o755); err != nil {
			return InitResult{}, err
		}
	}

	// Write starter files only if they don't exist
	starters := map[string]string{
		"csaw.yml":                       starterProfile,
		".csawignore":                    starterIgnore,
		"agents/base.md":                 starterAgent,
		"skills/commit-message/SKILL.md": starterSkill,
	}

	for relPath, content := range starters {
		fullPath := filepath.Join(absDir, relPath)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			if err := os.WriteFile(fullPath, []byte(content), 0o644); err != nil {
				return InitResult{}, err
			}
		}
	}

	gitDir := filepath.Join(absDir, ".git")
	if _, err := os.Stat(gitDir); os.IsNotExist(err) {
		if _, err := g.Run(ctx, absDir, "init"); err != nil {
			return InitResult{}, err
		}
	}

	return InitResult{Path: absDir, Name: name}, nil
}
