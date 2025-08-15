package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/loctree/internal/tree"
)

// Model represents the TUI application state
type Model struct {
	Root          *tree.DirectoryNode
	VisibleNodes  []*tree.DirectoryNode
	SelectedIndex int
	quitting      bool
}

// NewModel creates a new TUI model
func NewModel(root *tree.DirectoryNode) *Model {
	m := &Model{
		Root:          root,
		SelectedIndex: 0,
	}
	m.updateVisibleNodes()
	return m
}

// Init initializes the model (required by tea.Model)
func (m Model) Init() tea.Cmd {
	return nil
}

// Update handles messages and updates the model (required by tea.Model)
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			m.quitting = true
			return m, tea.Quit
			
		case "up", "k":
			if m.SelectedIndex > 0 {
				m.SelectedIndex--
			}
			
		case "down", "j":
			if m.SelectedIndex < len(m.VisibleNodes)-1 {
				m.SelectedIndex++
			}
		}
	}
	return m, nil
}

// View renders the model (required by tea.Model)
func (m Model) View() string {
	if m.quitting {
		return "Goodbye!\n"
	}
	
	return RenderTree(m.VisibleNodes, m.SelectedIndex)
}

// updateVisibleNodes rebuilds the list of visible nodes based on expanded state
func (m *Model) updateVisibleNodes() {
	m.VisibleNodes = nil
	m.addVisibleNodes(m.Root)
}

// addVisibleNodes recursively adds visible nodes to the list
func (m *Model) addVisibleNodes(node *tree.DirectoryNode) {
	m.VisibleNodes = append(m.VisibleNodes, node)
	
	if node.IsExpanded {
		for _, child := range node.Children {
			m.addVisibleNodes(child)
		}
	}
}