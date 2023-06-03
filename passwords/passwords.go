package passwords

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func Passwords(cmd *cobra.Command, args []string) {
	m := InitialModel()
	if _, err := tea.NewProgram(&m).Run(); err != nil {
		fmt.Println("Oh no!", err)
		os.Exit(1)
	}
}
