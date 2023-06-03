package passwords

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type KeyMap struct {
	quit     key.Binding
	cycle    key.Binding
	navigate key.Binding
}

type Model struct {
	Columns *Columns
	active  int

	keymap KeyMap
	help   help.Model
}

func InitialModel() Model {
	columns := Columns{
		Column{Letters: [6]rune{' ', ' ', ' ', ' ', ' ', ' '}, active: 0},
		Column{Letters: [6]rune{' ', ' ', ' ', ' ', ' ', ' '}, active: 0},
		Column{Letters: [6]rune{' ', ' ', ' ', ' ', ' ', ' '}, active: 0},
		Column{Letters: [6]rune{' ', ' ', ' ', ' ', ' ', ' '}, active: 0},
		Column{Letters: [6]rune{' ', ' ', ' ', ' ', ' ', ' '}, active: 0},
	}

	return Model{
		Columns: &columns,
		active:  0,
		keymap: KeyMap{
			quit: key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("ctrl+c", "quit"),
			),
			cycle: key.NewBinding(
				key.WithKeys("tab"),
				key.WithHelp("tab", "cycle"),
			),
			navigate: key.NewBinding(
				key.WithKeys("up", "down", "left", "right"),
				key.WithHelp("up/down/left/right", "navigate"),
			),
		},
		help: help.New(),
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.quit,
		m.keymap.cycle,
		m.keymap.navigate,
	})
}

func (m Model) View() string {
	matches := m.Solve()
	s := strings.Builder{}

	for _, word := range matches {
		s.WriteString(whitetext.Render(strings.ToUpper(word)))
		s.WriteRune(' ')
	}

	if len(matches) == 0 {
		s.WriteString("Type the letters that appear in each column.")
	}

	return lipgloss.JoinVertical(lipgloss.Center,
		outer_border.Render(lipgloss.JoinHorizontal(
			lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Center, title, m.Columns.View(m.active)),
			results.Render("Matches: "+greytext.Render(strconv.Itoa(len(matches)))+"\n\n"+s.String()),
		)),
		m.helpView(),
	)
}

func (m *Model) input(s rune) {
	active := m.Columns[m.active].active % 6
	m.Columns[m.active].Letters[active] = s
	m.Columns[m.active].active = (m.Columns[m.active].active + 1) % 6
}

func (m *Model) delete() {
	active := (m.Columns[m.active].active - 1) % 6
	if active < 0 {
		active = 5
	}

	m.Columns[m.active].Letters[active] = ' '
	m.Columns[m.active].active = active
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		s := msg.String()

		switch s {
		case "ctrl+c":
			return m, tea.Quit

		case "backspace":
			m.delete()

		case "tab":
			m.active = (m.active + 1) % 5

		case "left":
			active := (m.active - 1) % 5
			if active < 0 {
				active = 4
			}
			m.active = active

		case "right":
			m.active = (m.active + 1) % 5

		case "up":
			active := (m.Columns[m.active].active - 1) % 6
			if active < 0 {
				active = 5
			}
			m.Columns[m.active].active = active

		case "down":
			m.Columns[m.active].active = (m.Columns[m.active].active + 1) % 6

		default:
			if strings.Contains("abcdefghijklmnopqrstuvwxyz", s) {
				m.input(rune(s[0]))
			}
		}
	}

	return m, nil
}

func (m *Model) Solve() []string {
	potentialWords := WORDS

	for i, col := range m.Columns {
		potentialWords = filterWords(potentialWords, col.Letters, i)
	}

	return potentialWords
}

type Column struct {
	Letters [6]rune
	active  int
}

type Columns [5]Column

func (c Column) View(active bool) string {
	views := []string{}
	for i, letter := range c.Letters {
		if active && i == c.active%6 {
			views = append(views, GreyTile.Render(strings.ToUpper(string(letter))))
		} else {
			views = append(views, GreenTile.Render(strings.ToUpper(string(letter))))
		}
	}

	if active {
		return SelectedTile.Render(lipgloss.JoinVertical(lipgloss.Center, views...))
	} else {
		return UnselectedTile.Render(lipgloss.JoinVertical(lipgloss.Center, views...))
	}
}

func (cs Columns) View(active int) string {
	views := []string{}
	for i, c := range cs {

		views = append(views, c.View(i == active))
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, views...)
}
