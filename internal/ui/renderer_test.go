package ui

import (
	"strings"
	"testing"
	
	"github.com/user/loctree/internal/tree"
)

func TestRenderNode_Collapsed(t *testing.T) {
	node := tree.NewDirectoryNode("test", "/test")
	node.LOC = 100
	child := tree.NewDirectoryNode("child", "/test/child")
	node.AddChild(child)
	node.IsExpanded = false
	
	result := RenderNode(node, 0, false)
	
	if !strings.Contains(result, "▶") {
		t.Error("Expected collapsed indicator ▶")
	}
	if !strings.Contains(result, "100") {
		t.Error("Expected LOC count in output")
	}
	if !strings.Contains(result, "test") {
		t.Error("Expected node name in output")
	}
}

func TestRenderNode_Expanded(t *testing.T) {
	node := tree.NewDirectoryNode("test", "/test")
	node.LOC = 100
	child := tree.NewDirectoryNode("child", "/test/child")
	node.AddChild(child)
	node.IsExpanded = true
	
	result := RenderNode(node, 0, false)
	
	if !strings.Contains(result, "▼") {
		t.Error("Expected expanded indicator ▼")
	}
}

func TestRenderNode_WithIndentation(t *testing.T) {
	node := tree.NewDirectoryNode("test", "/test")
	node.LOC = 100
	
	result := RenderNode(node, 2, false)
	
	// Should have 4 spaces (2 levels * 2 spaces)
	if !strings.HasPrefix(result, "    ") {
		t.Error("Expected indentation of 4 spaces")
	}
}

func TestRenderTree(t *testing.T) {
	root := tree.NewDirectoryNode("root", "/root")
	root.LOC = 200
	root.IsExpanded = true
	
	child1 := tree.NewDirectoryNode("child1", "/root/child1")
	child1.LOC = 100
	child2 := tree.NewDirectoryNode("child2", "/root/child2")
	child2.LOC = 50
	
	root.AddChild(child1)
	root.AddChild(child2)
	
	visibleNodes := []*tree.DirectoryNode{root, child1, child2}
	selectedIndex := 0
	
	result := RenderTree(visibleNodes, selectedIndex)
	
	lines := strings.Split(result, "\n")
	if len(lines) < 3 {
		t.Errorf("Expected at least 3 lines, got %d", len(lines))
	}
	
	// First line should be highlighted (selected)
	// This will depend on the styling implementation
}

func TestRenderNode_NoChildren(t *testing.T) {
	node := tree.NewDirectoryNode("leaf", "/leaf")
	node.LOC = 50
	
	result := RenderNode(node, 0, false)
	
	// Should not have any indicator
	if strings.Contains(result, "▶") || strings.Contains(result, "▼") {
		t.Error("Leaf node should not have expand/collapse indicator")
	}
}