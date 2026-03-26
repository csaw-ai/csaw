package drift

import (
	"os"
	"path/filepath"

	"github.com/csaw-ai/csaw/internal/runtime"
	"github.com/csaw-ai/csaw/internal/workspace"
)

const (
	IssueMissingSource = "missing-source"
	IssueDriftedLink   = "drifted-link"
	IssueMissingLink   = "missing-link"
	IssueReplacedLink  = "replaced-link"
)

type Status struct {
	Healthy        bool
	RelativePath   string
	FullPath       string
	SourceName     string
	ExpectedSource string
	ActualTarget   string
	ResolvedTarget string
	Issue          string
}

func InspectLinks(links []workspace.MountedLink) []Status {
	statuses := make([]Status, 0, len(links))
	for _, link := range links {
		status := Status{
			Healthy:        true,
			RelativePath:   link.RelativePath,
			FullPath:       link.FullPath,
			ResolvedTarget: link.ResolvedTarget,
		}

		if _, err := os.Stat(link.ResolvedTarget); err != nil {
			status.Healthy = false
			status.Issue = IssueMissingSource
		}

		statuses = append(statuses, status)
	}

	return statuses
}

func InspectMountState(projectRoot string, state workspace.MountState) []Status {
	statuses := make([]Status, 0, len(state.Entries))
	for _, entry := range state.Entries {
		fullPath := filepath.Join(projectRoot, filepath.FromSlash(entry.RelativePath))
		status := Status{
			Healthy:        true,
			RelativePath:   entry.RelativePath,
			FullPath:       fullPath,
			SourceName:     entry.SourceName,
			ExpectedSource: entry.SourcePath,
		}

		if _, err := os.Stat(entry.SourcePath); err != nil {
			status.Healthy = false
			status.Issue = IssueMissingSource
		}

		info, err := os.Lstat(fullPath)
		if err != nil {
			if os.IsNotExist(err) {
				status.Healthy = false
				status.Issue = IssueMissingLink
			}
			statuses = append(statuses, status)
			continue
		}

		if info.Mode()&os.ModeSymlink == 0 {
			status.Healthy = false
			status.Issue = IssueReplacedLink
			statuses = append(statuses, status)
			continue
		}

		actualTarget, err := os.Readlink(fullPath)
		if err != nil {
			status.Healthy = false
			statuses = append(statuses, status)
			continue
		}
		status.ActualTarget = actualTarget
		resolvedTarget := actualTarget
		if !filepath.IsAbs(resolvedTarget) {
			resolvedTarget = filepath.Join(filepath.Dir(fullPath), resolvedTarget)
		}
		status.ResolvedTarget = runtime.NormalizeFSPath(resolvedTarget)

		if status.Issue == IssueMissingSource {
			statuses = append(statuses, status)
			continue
		}

		if !runtime.PathsEqual(status.ResolvedTarget, entry.SourcePath) {
			status.Healthy = false
			status.Issue = IssueDriftedLink
		}

		statuses = append(statuses, status)
	}

	return statuses
}
