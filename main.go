package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
)

var direct = flag.Bool("direct", false, "Editor directly, not expose SSH server")

func main() {
	flag.Parse()
	m := model{}
	if *direct {
		if _, err := tea.NewProgram(m).Run(); err != nil {
			fatal(err)
		}
		os.Exit(0)
	}

	// Expose over SSH server
	addr := "localhost:23553"
	slog.Info(fmt.Sprintf("listening on %s", addr))
	server, err := wish.NewServer(
		wish.WithAddress(addr),
		wish.WithMiddleware(myCustomBubbleteaMiddleware()),
	)
	if err != nil {
		fatal(err)
	}

	if err := server.ListenAndServe(); err != nil {
		fatal(err)
	}
}

func fatal(err any) {
	slog.Error(fmt.Sprintf("Error: %s", err))
	os.Exit(1)
}
