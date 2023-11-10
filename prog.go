package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type editorFinishedMsg struct{ err error }

func openEditor() tea.Cmd {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	slog.Info(fmt.Sprintf("opening editor %q", editor))
	c := exec.Command(editor) //nolint:gosec
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return editorFinishedMsg{err}
	})
}

type model struct{}

func (m model) Init() tea.Cmd {
	return openEditor()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case editorFinishedMsg:
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	return ""
}
