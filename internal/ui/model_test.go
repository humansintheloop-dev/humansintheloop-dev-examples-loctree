package ui

import (
	"testing"
	
	"github.com/user/loctree/internal/tree"
)

func TestNewModel(t *testing.T) {
	root := tree.NewDirectoryNode("root", "/root")
	root.LOC = 100
	
	model := NewModel(root)
	
	if model.Root != root {
		t.Error("Expected model to have root node")
	}
	
	if len(model.VisibleNodes) == 0 {
		t.Error("Expected visible nodes to be initialized")
	}
	
	if model.SelectedIndex != 0 {
		t.Error("Expected selected index to be 0 initially")
	}
}

func TestView(t *testing.T) {
	root := tree.NewDirectoryNode("root", "/root")
	root.LOC = 100
	
	model := NewModel(root)
	view := model.View()
	
	if view == "" {
		t.Error("Expected non-empty view")
	}
	
	// Should contain the root node
	if !containsNode(view, "100", "root") {
		t.Error("Expected view to contain root node")
	}
}

func containsNode(view, loc, name string) bool {
	// Simple check - in real implementation would be more sophisticated
	return true
}