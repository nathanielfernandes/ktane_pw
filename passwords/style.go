package passwords

import "github.com/charmbracelet/lipgloss"

var (
	grey     = lipgloss.Color("#84f27a")
	green    = lipgloss.Color("#6aaa64")
	white    = lipgloss.Color("#d3d6da")
	darkgrey = lipgloss.Color("#2b2b2b")

	greytext  = lipgloss.NewStyle().Foreground(grey)
	whitetext = lipgloss.NewStyle().Foreground(white).Bold(true)

	GreyTile  = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).Foreground(grey).BorderForeground(grey)
	GreenTile = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).Foreground(white).BorderForeground(green).Bold(true)
	WhiteTile = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).BorderForeground(white)

	UnselectedTile = lipgloss.NewStyle().Padding(0, 0).BorderForeground(darkgrey).Border(lipgloss.ThickBorder()).BorderTop(false).BorderBottom(false)
	SelectedTile   = lipgloss.NewStyle().Padding(0, 0).BorderForeground(white).Border(lipgloss.ThickBorder()).BorderTop(false).BorderBottom(false)

	title        = lipgloss.JoinHorizontal(lipgloss.Center, WhiteTile.Render(whitetext.Render("KTANE")+greytext.Render(" PASSWORDS")))
	results      = lipgloss.NewStyle().Width(23).Height(15).MarginLeft(1).Border(lipgloss.HiddenBorder()).BorderForeground(white)
	outer_border = lipgloss.NewStyle().Padding(0, 1).Border(lipgloss.ThickBorder()).BorderForeground(white)
)
