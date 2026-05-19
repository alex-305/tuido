package screens

import (
	"fmt"
	"os"

	"github.com/alex-305/tuido/internal/context"
	"github.com/alex-305/tuido/internal/ui/screens/homescreen"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	Init() tea.Cmd
	Update(msg tea.Msg, width, height int) (Screen, tea.Cmd)
	View(width, height int) string
}

type Model struct {
	width  int
	height int

	history []Screen
	current Screen

	ctx context.AppContext
}

func NewModel() *Model {

	var initialScreen Screen

	ctx := context.AppContext{}

	initialScreen = homescreen.NewHomeScreen(ctx)

	return &Model{
		current: initialScreen,
		history: []Screen{},
		ctx:     ctx,
	}
}

func (m *Model) Init() tea.Cmd {
	if m.current != nil {
		return m.current.Init()
	}
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	h, c, ok := m.handleMessages(msg)
	if ok {
		return h, c
	}

	keyMsg, isKeyMsg := msg.(tea.KeyMsg)
	if isKeyMsg {
		h, c, ok := m.handleKeyMsg(keyMsg)
		if ok {
			return h, c
		}
	}

	return m.updateScreen(msg)

}

func (m *Model) updateScreen(msg tea.Msg) (tea.Model, tea.Cmd) {
	screen, cmd := m.current.Update(msg, m.width, m.height)
	m.current = screen
	return m, cmd
}

func (m *Model) View() string {
	return m.current.View(m.width, m.height)
}
