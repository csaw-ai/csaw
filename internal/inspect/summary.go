package inspect

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"

	"github.com/csaw-ai/csaw/internal/drift"
	"github.com/csaw-ai/csaw/internal/runtime"
	"github.com/csaw-ai/csaw/internal/sources"
	"github.com/csaw-ai/csaw/internal/workspace"
)

type Summary struct {
	ProjectRoot string
	Paths       runtime.Paths
	Sources     []sources.Source
	Mounted     []drift.Status
}

func BuildSummary(ctx context.Context, projectRoot string, paths runtime.Paths, manager sources.Manager) (Summary, error) {
	_ = ctx

	cfg, err := manager.Load()
	if err != nil {
		return Summary{}, err
	}

	state, err := workspace.ReadMountState(projectRoot)
	if err != nil {
		return Summary{}, err
	}

	var mounted []drift.Status
	if len(state.Entries) > 0 {
		mounted = drift.InspectMountState(projectRoot, state)
	} else {
		links, err := workspace.FindMountedLinks(projectRoot, paths.Root)
		if err != nil {
			return Summary{}, err
		}
		mounted = drift.InspectLinks(links)
	}

	return Summary{
		ProjectRoot: projectRoot,
		Paths:       paths,
		Sources:     cfg.Sources,
		Mounted:     mounted,
	}, nil
}

func RenderSummary(summary Summary) string {
	title := lipgloss.NewStyle().Bold(true).Render("csaw inspect")
	key := lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	okStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	warnStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("11"))

	var builder strings.Builder
	builder.WriteString(title)
	builder.WriteString("\n\n")
	builder.WriteString(fmt.Sprintf("%s %s\n", key.Render("project:"), summary.ProjectRoot))
	builder.WriteString(fmt.Sprintf("%s %s\n", key.Render("csaw home:"), summary.Paths.Root))
	builder.WriteString(fmt.Sprintf("%s %d\n", key.Render("sources:"), len(summary.Sources)))
	builder.WriteString(fmt.Sprintf("%s %d\n", key.Render("mounted links:"), len(summary.Mounted)))

	if len(summary.Sources) > 0 {
		builder.WriteString("\nConfigured sources:\n")
		for _, source := range summary.Sources {
			builder.WriteString(fmt.Sprintf("- %s (%s) -> %s\n", source.Name, source.Kind, source.CheckoutPath(summary.Paths)))
		}
	}

	if len(summary.Mounted) > 0 {
		builder.WriteString("\nMounted links:\n")
		for _, status := range summary.Mounted {
			label := okStyle.Render("healthy")
			if !status.Healthy {
				label = warnStyle.Render(status.Issue)
			}
			if status.SourceName != "" {
				builder.WriteString(fmt.Sprintf("- %s %s (%s) -> %s\n", label, status.RelativePath, status.SourceName, status.ExpectedSource))
				continue
			}
			builder.WriteString(fmt.Sprintf("- %s %s -> %s\n", label, status.RelativePath, status.ResolvedTarget))
		}
	}

	return builder.String()
}

func RenderSourceDetails(source sources.Source, paths runtime.Paths) (string, error) {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("name:\t%s\n", source.Name))
	builder.WriteString(fmt.Sprintf("kind:\t%s\n", source.Kind))
	if source.URL != "" {
		builder.WriteString(fmt.Sprintf("url:\t%s\n", source.URL))
	}
	if source.Path != "" {
		builder.WriteString(fmt.Sprintf("path:\t%s\n", source.Path))
	}
	builder.WriteString(fmt.Sprintf("checkout:\t%s\n", source.CheckoutPath(paths)))
	return builder.String(), nil
}

func RenderMarkdownPreview(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	renderer, err := glamour.NewTermRenderer(glamour.WithAutoStyle())
	if err != nil {
		return "", err
	}

	return renderer.Render(string(content))
}
