package utils

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Text conatins all the functions related to text manipulations
type Text struct{}

// P is the amount of padding required
type P struct {
	Left   int
	Top    int
	Bottom int
}

// Style is a struct that contains the styling properties
type Style struct {
	Color   lipgloss.TerminalColor
	Padding P
	Align   lipgloss.Position
	Bold    bool
}

// H heading
func (Text) H(style Style, strs ...string,
) string {
	return lipgloss.NewStyle().
		Bold(style.Bold).
		Foreground(style.Color).
		PaddingTop(style.Padding.Top).
		PaddingLeft(style.Padding.Left).
		PaddingBottom(style.Padding.Bottom).
		Align(style.Align).
		Render(strs...)
}

// P paragraph
func (Text) P(style Style, strs ...string,
) string {
	return lipgloss.NewStyle().
		Bold(style.Bold).
		Foreground(style.Color).
		PaddingTop(style.Padding.Top).
		PaddingLeft(style.Padding.Left).
		PaddingBottom(style.Padding.Bottom).
		Align(style.Align).
		Render(strs...)
}

// Error is a function that is used to display errors to the standered output
func (Text) Error(strs ...string) {
	fmt.Println(
		Text{}.P(Style{
			Bold:  false,
			Color: lipgloss.Color("#D72023"),
			Padding: P{
				Left:   1,
				Top:    1,
				Bottom: 1,
			},
			Align: lipgloss.Left,
		}, strs...),
	)
}
