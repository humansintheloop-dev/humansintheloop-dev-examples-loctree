package tree

import (
	"testing"
)

func TestNewDirectoryNode(t *testing.T) {
	node := NewDirectoryNode("test", "/path/to/test")
	
	if node.Name != "test" {
		t.Errorf("Expected name 'test', got '%s'", node.Name)
	}
	
	if node.Path != "/path/to/test" {
		t.Errorf("Expected path '/path/to/test', got '%s'", node.Path)
	}
	
	if node.LOC != 0 {
		t.Errorf("Expected LOC 0, got %d", node.LOC)
	}
	
	if node.IsExpanded {
		t.Error("Expected IsExpanded to be false initially")
	}
	
	if len(node.Children) != 0 {
		t.Errorf("Expected no children initially, got %d", len(node.Children))
	}
	
	if node.Parent != nil {
		t.Error("Expected Parent to be nil for root node")
	}
}

func TestAddChild(t *testing.T) {
	parent := NewDirectoryNode("parent", "/parent")
	child := NewDirectoryNode("child", "/parent/child")
	
	parent.AddChild(child)
	
	if len(parent.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(parent.Children))
	}
	
	if parent.Children[0] != child {
		t.Error("Child not added correctly")
	}
	
	if child.Parent != parent {
		t.Error("Parent-child relationship not established")
	}
}

func TestCalculateLOC(t *testing.T) {
	root := NewDirectoryNode("root", "/root")
	root.FileLOC = 10 // Files in root directory
	
	child1 := NewDirectoryNode("child1", "/root/child1")
	child1.FileLOC = 20
	
	child2 := NewDirectoryNode("child2", "/root/child2")
	child2.FileLOC = 30
	
	grandchild := NewDirectoryNode("grandchild", "/root/child1/grandchild")
	grandchild.FileLOC = 15
	
	root.AddChild(child1)
	root.AddChild(child2)
	child1.AddChild(grandchild)
	
	root.CalculateLOC()
	
	// grandchild should have 15
	if grandchild.LOC != 15 {
		t.Errorf("Expected grandchild LOC 15, got %d", grandchild.LOC)
	}
	
	// child1 should have 20 + 15 = 35
	if child1.LOC != 35 {
		t.Errorf("Expected child1 LOC 35, got %d", child1.LOC)
	}
	
	// child2 should have 30
	if child2.LOC != 30 {
		t.Errorf("Expected child2 LOC 30, got %d", child2.LOC)
	}
	
	// root should have 10 + 35 + 30 = 75
	if root.LOC != 75 {
		t.Errorf("Expected root LOC 75, got %d", root.LOC)
	}
}