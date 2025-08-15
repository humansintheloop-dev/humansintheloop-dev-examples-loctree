package ui

import (
	"fmt"
	"strings"
	
	"github.com/charmbracelet/lipgloss"
	"github.com/user/loctree/internal/tree"
)

var (
	normalStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("252"))
	
	selectedStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true)
	
	indicatorStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
	
	locStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("214"))
)

// RenderNode renders a single node with proper formatting
func RenderNode(node *tree.DirectoryNode, depth int, selected bool) string {
	indent := strings.Repeat("  ", depth)
	
	// Determine indicator
	indicator := ""
	if len(node.Children) > 0 {
		if node.IsExpanded {
			indicator = "▼ "
		} else {
			indicator = "▶ "
		}
	}
	
	// Format the line
	line := fmt.Sprintf("%s%s%d %s", indent, indicator, node.LOC, node.Name)
	
	// Apply style based on selection
	if selected {
		return selectedStyle.Render(line)
	}
	return normalStyle.Render(line)
}

// RenderTree renders the entire visible tree
func RenderTree(visibleNodes []*tree.DirectoryNode, selectedIndex int) string {
	var lines []string
	
	for i, node := range visibleNodes {
		selected := i == selectedIndex
		depth := getNodeDepth(node)
		line := RenderNode(node, depth, selected)
		lines = append(lines, line)
	}
	
	return strings.Join(lines, "\n")
}

// getNodeDepth calculates the depth of a node in the tree
func getNodeDepth(node *tree.DirectoryNode) int {
	depth := 0
	current := node
	for current.Parent != nil {
		depth++
		current = current.Parent
	}
	return depth
}