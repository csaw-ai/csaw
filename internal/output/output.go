package output

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Semantic colors
	colorInfo    = lipgloss.Color("12") // blue
	colorSuccess = lipgloss.Color("10") // green
	colorWarn    = lipgloss.Color("11") // yellow
	colorError   = lipgloss.Color("9")  // red
	colorMuted   = lipgloss.Color("8")  // gray
	colorAccent  = lipgloss.Color("14") // cyan

	// Text styles
	infoStyle    = lipgloss.NewStyle().Foreground(colorInfo)
	successStyle = lipgloss.NewStyle().Foreground(colorSuccess)
	warnStyle    = lipgloss.NewStyle().Foreground(colorWarn)
	errorStyle   = lipgloss.NewStyle().Foreground(colorError)
	mutedStyle   = lipgloss.NewStyle().Foreground(colorMuted)
	accentStyle  = lipgloss.NewStyle().Foreground(colorAccent)
	boldStyle    = lipgloss.NewStyle().Bold(true)
	labelStyle   = lipgloss.NewStyle().Foreground(colorMuted).Width(14)

	// Symbols
	SymbolOK   = successStyle.Render("✔")
	SymbolWarn = warnStyle.Render("!")
	SymbolErr  = errorStyle.Render("✖")
	SymbolDot  = mutedStyle.Render("·")
	SymbolArr  = mutedStyle.Render("→")
)

func Infof(format string, args ...any) {
	fmt.Println(infoStyle.Render(fmt.Sprintf(format, args...)))
}

func Successf(format string, args ...any) {
	fmt.Printf("%s %s\n", SymbolOK, fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	fmt.Printf("%s %s\n", SymbolWarn, warnStyle.Render(fmt.Sprintf(format, args...)))
}

func Errorf(format string, args ...any) {
	fmt.Printf("%s %s\n", SymbolErr, errorStyle.Render(fmt.Sprintf(format, args...)))
}

// Header prints a bold section header.
func Header(text string) {
	fmt.Println(boldStyle.Render(text))
}

// Label prints a key-value pair with aligned labels.
func Label(key, value string) {
	fmt.Printf("%s %s\n", labelStyle.Render(key), value)
}

// Muted prints dimmed secondary text.
func Muted(text string) {
	fmt.Println(mutedStyle.Render(text))
}

// Accent returns text in the accent color (for highlighting names, paths).
func Accent(text string) string {
	return accentStyle.Render(text)
}

// Bold returns bolded text.
func Bold(text string) string {
	return boldStyle.Render(text)
}

// Success returns green text.
func Success(text string) string {
	return successStyle.Render(text)
}

// Warn returns yellow text.
func Warn(text string) string {
	return warnStyle.Render(text)
}

// Faint returns muted text.
func Faint(text string) string {
	return mutedStyle.Render(text)
}

// Divider prints a horizontal rule.
func Divider() {
	fmt.Println(mutedStyle.Render(strings.Repeat("─", 50)))
}
