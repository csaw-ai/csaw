// Package linkmode handles cross-platform file linking.
//
// On Unix systems (and Windows with Developer Mode enabled), csaw creates
// symlinks. On Windows without symlink privileges, it falls back to hardlinks
// which preserve the same live-update behavior for individual files.
package linkmode

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// Mode represents the linking strategy in use.
type Mode int

const (
	Symlink  Mode = iota // os.Symlink — default on Unix, Windows with Developer Mode
	Hardlink             // os.Link — Windows fallback, no privileges required
)

func (m Mode) String() string {
	switch m {
	case Symlink:
		return "symlink"
	case Hardlink:
		return "hardlink"
	default:
		return "unknown"
	}
}

// Detect probes the system to determine which linking mode is available.
// It creates a temporary symlink to test whether the OS permits it.
// On non-Windows systems this always returns Symlink.
func Detect() Mode {
	if runtime.GOOS != "windows" {
		return Symlink
	}

	dir, err := os.MkdirTemp("", "csaw-probe-*")
	if err != nil {
		return Hardlink
	}
	defer os.RemoveAll(dir)

	target := filepath.Join(dir, "target")
	link := filepath.Join(dir, "link")
	if err := os.WriteFile(target, []byte("probe"), 0o644); err != nil {
		return Hardlink
	}
	if err := os.Symlink(target, link); err != nil {
		return Hardlink
	}
	return Symlink
}

// Create creates a link from source to target using the given mode.
// For Symlink mode, source is the symlink target (what it points to).
// For Hardlink mode, source is the existing file to link to.
func Create(mode Mode, source, target string) error {
	switch mode {
	case Symlink:
		return os.Symlink(source, target)
	case Hardlink:
		if err := os.Link(source, target); err != nil {
			return fmt.Errorf("hardlink failed (source and target must be on the same volume): %w", err)
		}
		return nil
	default:
		return fmt.Errorf("unsupported link mode: %v", mode)
	}
}

// IsLink reports whether the file at path is a csaw-managed link.
// For symlinks it checks ModeSymlink; for hardlinks it compares file identity
// with the expected source using os.SameFile.
func IsLink(mode Mode, path, expectedSource string) bool {
	switch mode {
	case Symlink:
		info, err := os.Lstat(path)
		if err != nil {
			return false
		}
		return info.Mode()&os.ModeSymlink != 0
	case Hardlink:
		pathInfo, err := os.Stat(path)
		if err != nil {
			return false
		}
		sourceInfo, err := os.Stat(expectedSource)
		if err != nil {
			return false
		}
		return os.SameFile(pathInfo, sourceInfo)
	default:
		return false
	}
}

// ReadTarget returns the resolved source path for a link.
// For symlinks it uses os.Readlink. For hardlinks, the source path cannot be
// derived from the filesystem — the caller must supply it from mount state.
func ReadTarget(mode Mode, path string) (string, error) {
	switch mode {
	case Symlink:
		target, err := os.Readlink(path)
		if err != nil {
			return "", err
		}
		if !filepath.IsAbs(target) {
			target = filepath.Join(filepath.Dir(path), target)
		}
		return target, nil
	case Hardlink:
		return "", fmt.Errorf("cannot read target of hardlink; use mount state")
	default:
		return "", fmt.Errorf("unsupported link mode: %v", mode)
	}
}

// Verify checks whether the link at path correctly points to expectedSource.
// For symlinks it resolves and compares paths. For hardlinks it uses os.SameFile.
func Verify(mode Mode, path, expectedSource string, pathsEqual func(a, b string) bool) (healthy bool, actualTarget string) {
	switch mode {
	case Symlink:
		resolved, err := ReadTarget(mode, path)
		if err != nil {
			return false, ""
		}
		return pathsEqual(resolved, expectedSource), resolved
	case Hardlink:
		pathInfo, err := os.Stat(path)
		if err != nil {
			return false, ""
		}
		sourceInfo, err := os.Stat(expectedSource)
		if err != nil {
			return false, ""
		}
		if os.SameFile(pathInfo, sourceInfo) {
			return true, expectedSource
		}
		return false, path
	default:
		return false, ""
	}
}

// UnavailableError returns a user-facing error message when neither symlinks
// nor hardlinks are available (e.g., cross-volume on Windows without Developer Mode).
func UnavailableError(source, target string) error {
	return fmt.Errorf(
		"cannot link %s → %s: symlinks require Developer Mode and hardlinks require the same drive\n\n"+
			"To fix this, either:\n"+
			"  1. Enable Developer Mode: Settings → Privacy & Security → For developers\n"+
			"  2. Ensure your project and ~/.csaw are on the same drive",
		target, source,
	)
}
