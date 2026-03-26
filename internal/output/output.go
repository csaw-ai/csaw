package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	infoStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	successStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("10"))
	warnStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
)

func Infof(format string, args ...any) {
	fmt.Println(infoStyle.Render(fmt.Sprintf(format, args...)))
}

func Successf(format string, args ...any) {
	fmt.Println(successStyle.Render(fmt.Sprintf(format, args...)))
}

func Warnf(format string, args ...any) {
	fmt.Println(warnStyle.Render(fmt.Sprintf(format, args...)))
}
