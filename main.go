package main

import (
	"fmt"
	"os"

	"github.com/alex-305/tuido/internal/context"
	"github.com/alex-305/tuido/internal/storage"
	"github.com/alex-305/tuido/internal/ui/screens"
	"github.com/alex-305/tuido/internal/ui/screens/homescreen"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	appContext := context.AppContext{}
	store, err := storage.NewStore()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize database: %v\n", err)
		os.Exit(1)
	}

	appContext.Store = store

	p := tea.NewProgram(screens.NewModel(homescreen.NewHomeScreen(appContext)), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
