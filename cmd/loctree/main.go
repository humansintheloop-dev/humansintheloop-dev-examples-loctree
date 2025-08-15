package main

import (
	"fmt"
	"os"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/loctree/internal/cli"
	"github.com/user/loctree/internal/ui"
)

func main() {
	// Parse command-line arguments
	path, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	
	// Validate the path
	err = cli.ValidatePath(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	
	// Create and run TUI with loading screen
	model := ui.NewLoadingModel(path)
	p := tea.NewProgram(model, tea.WithAltScreen())
	
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running TUI: %v\n", err)
		os.Exit(1)
	}
}