package git

import (
	"bytes"
	"context"
	"os/exec"
	"strings"
)

type Git interface {
	Run(ctx context.Context, cwd string, args ...string) (string, error)
}

type ExecGit struct{}

func (ExecGit) Run(ctx context.Context, cwd string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	if cwd != "" {
		cmd.Dir = cwd
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	output := strings.TrimSpace(stdout.String())
	if err != nil {
		message := strings.TrimSpace(stderr.String())
		if message == "" {
			message = err.Error()
		}
		return output, &execError{message: message}
	}

	return output, nil
}

type execError struct {
	message string
}

func (e *execError) Error() string {
	return e.message
}
