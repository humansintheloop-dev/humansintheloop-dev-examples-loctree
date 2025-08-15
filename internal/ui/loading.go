package ui

import (
	"fmt"
	"time"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/loctree/internal/tree"
)

// LoadingModel shows a loading indicator while scanning
type LoadingModel struct {
	path     string
	done     bool
	root     *tree.DirectoryNode
	err      error
	spinner  int
}

var spinnerFrames = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}

// NewLoadingModel creates a new loading model
func NewLoadingModel(path string) *LoadingModel {
	return &LoadingModel{
		path: path,
	}
}

// Init starts the tree building process
func (m LoadingModel) Init() tea.Cmd {
	return tea.Batch(
		buildTreeCmd(m.path),
		tickCmd(),
	)
}

// Update handles messages
func (m LoadingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case treeBuiltMsg:
		m.done = true
		m.root = msg.root
		m.err = msg.err
		if m.err != nil {
			return m, tea.Quit
		}
		// Switch to main model
		return NewModel(m.root), nil
		
	case tickMsg:
		m.spinner = (m.spinner + 1) % len(spinnerFrames)
		if !m.done {
			return m, tickCmd()
		}
		
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" || msg.String() == "q" {
			return m, tea.Quit
		}
	}
	
	return m, nil
}

// View renders the loading screen
func (m LoadingModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("Error: %v\n", m.err)
	}
	
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)
	
	spinner := spinnerFrames[m.spinner]
	return style.Render(fmt.Sprintf("%s Scanning: %s", spinner, m.path))
}

// Messages
type treeBuiltMsg struct {
	root *tree.DirectoryNode
	err  error
}

type tickMsg struct{}

// Commands
func buildTreeCmd(path string) tea.Cmd {
	return func() tea.Msg {
		root, err := tree.BuildTree(path)
		return treeBuiltMsg{root: root, err: err}
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(100, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}