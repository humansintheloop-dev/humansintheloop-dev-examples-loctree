package tree

import (
	"testing"
)

func TestToggleExpanded_Collapsed(t *testing.T) {
	node := NewDirectoryNode("test", "/test")
	child := NewDirectoryNode("child", "/test/child")
	node.AddChild(child)
	node.IsExpanded = false
	
	node.ToggleExpanded()
	
	if !node.IsExpanded {
		t.Error("Expected node to be expanded after toggle")
	}
}

func TestToggleExpanded_Expanded(t *testing.T) {
	node := NewDirectoryNode("test", "/test")
	child := NewDirectoryNode("child", "/test/child")
	node.AddChild(child)
	node.IsExpanded = true
	
	node.ToggleExpanded()
	
	if node.IsExpanded {
		t.Error("Expected node to be collapsed after toggle")
	}
}

func TestGetVisibleNodes_AllCollapsed(t *testing.T) {
	root := NewDirectoryNode("root", "/root")
	child1 := NewDirectoryNode("child1", "/root/child1")
	child2 := NewDirectoryNode("child2", "/root/child2")
	root.AddChild(child1)
	root.AddChild(child2)
	root.IsExpanded = false
	
	visible := GetVisibleNodes(root)
	
	if len(visible) != 1 {
		t.Errorf("Expected 1 visible node (root only), got %d", len(visible))
	}
}

func TestGetVisibleNodes_SomeExpanded(t *testing.T) {
	root := NewDirectoryNode("root", "/root")
	child1 := NewDirectoryNode("child1", "/root/child1")
	child2 := NewDirectoryNode("child2", "/root/child2")
	grandchild := NewDirectoryNode("grandchild", "/root/child1/grandchild")
	
	root.AddChild(child1)
	root.AddChild(child2)
	child1.AddChild(grandchild)
	
	root.IsExpanded = true
	child1.IsExpanded = true
	
	visible := GetVisibleNodes(root)
	
	// Should have root, child1, grandchild, child2
	if len(visible) != 4 {
		t.Errorf("Expected 4 visible nodes, got %d", len(visible))
	}
}