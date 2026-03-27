package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// PickerItem represents a selectable item in the profile picker.
type PickerItem struct {
	Name        string
	Description string
	Detail      string // e.g., "extends: base, 5 files"
}

// PickerResult holds the outcome of the picker interaction.
type PickerResult struct {
	Selected string
	Aborted  bool
}

type pickerModel struct {
	items     []PickerItem
	filtered  []PickerItem
	cursor    int
	textInput textinput.Model
	result    PickerResult
	width     int
	height    int
}

var (
	selectedStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("14")).Bold(true)
	unselectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	descStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	detailStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Italic(true)
	promptStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	helpStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)

// RunPicker shows an interactive fuzzy profile picker and returns the selection.
func RunPicker(items []PickerItem, prompt string) (PickerResult, error) {
	if len(items) == 0 {
		return PickerResult{Aborted: true}, nil
	}

	ti := textinput.New()
	ti.Placeholder = "type to filter..."
	ti.Focus()
	ti.CharLimit = 64

	model := pickerModel{
		items:     items,
		filtered:  items,
		textInput: ti,
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return PickerResult{}, err
	}

	return finalModel.(pickerModel).result, nil
}

func (m pickerModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m pickerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			m.result = PickerResult{Aborted: true}
			return m, tea.Quit

		case "enter":
			if len(m.filtered) > 0 && m.cursor < len(m.filtered) {
				m.result = PickerResult{Selected: m.filtered[m.cursor].Name}
			} else {
				m.result = PickerResult{Aborted: true}
			}
			return m, tea.Quit

		case "up", "ctrl+p":
			if m.cursor > 0 {
				m.cursor--
			}
			return m, nil

		case "down", "ctrl+n":
			if m.cursor < len(m.filtered)-1 {
				m.cursor++
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)

	// Re-filter on input change
	query := strings.ToLower(m.textInput.Value())
	if query == "" {
		m.filtered = m.items
	} else {
		m.filtered = m.filtered[:0]
		for _, item := range m.items {
			if fuzzyMatch(query, item.Name) || fuzzyMatch(query, item.Description) {
				m.filtered = append(m.filtered, item)
			}
		}
	}

	// Keep cursor in bounds
	if m.cursor >= len(m.filtered) {
		m.cursor = max(0, len(m.filtered)-1)
	}

	return m, cmd
}

func (m pickerModel) View() string {
	var b strings.Builder

	b.WriteString(promptStyle.Render("Select a profile"))
	b.WriteString("\n")
	b.WriteString(m.textInput.View())
	b.WriteString("\n\n")

	// Show items (limit visible to terminal height)
	maxVisible := 12
	if m.height > 0 {
		maxVisible = max(5, m.height-6)
	}

	start := 0
	if m.cursor >= maxVisible {
		start = m.cursor - maxVisible + 1
	}
	end := min(start+maxVisible, len(m.filtered))

	for i := start; i < end; i++ {
		item := m.filtered[i]
		cursor := "  "
		nameStyle := unselectedStyle
		if i == m.cursor {
			cursor = selectedStyle.Render("▸ ")
			nameStyle = selectedStyle
		}

		line := fmt.Sprintf("%s%s", cursor, nameStyle.Render(item.Name))
		if item.Description != "" {
			line += " " + descStyle.Render(item.Description)
		}
		b.WriteString(line)
		if i == m.cursor && item.Detail != "" {
			b.WriteString("\n    " + detailStyle.Render(item.Detail))
		}
		b.WriteString("\n")
	}

	if len(m.filtered) == 0 {
		b.WriteString(descStyle.Render("  no matching profiles"))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(helpStyle.Render("↑↓ navigate · enter select · esc cancel"))

	return b.String()
}

// fuzzyMatch does a simple substring match. Could be upgraded to
// real fuzzy matching later.
func fuzzyMatch(query, target string) bool {
	return strings.Contains(strings.ToLower(target), query)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
