package main

import (
	"fmt"
	"os"

	"github.com/alex-305/tuido/internal/ui/screens"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(screens.NewModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
