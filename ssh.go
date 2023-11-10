package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/muesli/termenv"
)

func myCustomBubbleteaMiddleware() wish.Middleware {
	newProg := func(m tea.Model, opts ...tea.ProgramOption) *tea.Program {
		p := tea.NewProgram(m, opts...)
		return p
	}
	teaHandler := func(s ssh.Session) *tea.Program {
		_, _, active := s.Pty()
		if !active {
			wish.Fatalln(s, "no active terminal, skipping")
			return nil
		}
		return newProg(model{},
			tea.WithInput(s),
			tea.WithOutput(s),
			tea.WithAltScreen(),
			tea.WithInputTTY(),
			tea.WithContext(s.Context()),
		)
	}
	return bm.MiddlewareWithProgramHandler(teaHandler, termenv.Ascii)
}
