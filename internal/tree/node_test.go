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

func TestSortChildren_ByLOC(t *testing.T) {
	root := NewDirectoryNode("root", "/root")
	
	child1 := NewDirectoryNode("small", "/root/small")
	child1.LOC = 10
	
	child2 := NewDirectoryNode("large", "/root/large")
	child2.LOC = 50
	
	child3 := NewDirectoryNode("medium", "/root/medium")
	child3.LOC = 25
	
	root.AddChild(child1)
	root.AddChild(child2)
	root.AddChild(child3)
	
	root.SortChildren()
	
	// Should be sorted by LOC descending: large (50), medium (25), small (10)
	if root.Children[0].Name != "large" {
		t.Errorf("Expected first child to be 'large', got '%s'", root.Children[0].Name)
	}
	if root.Children[1].Name != "medium" {
		t.Errorf("Expected second child to be 'medium', got '%s'", root.Children[1].Name)
	}
	if root.Children[2].Name != "small" {
		t.Errorf("Expected third child to be 'small', got '%s'", root.Children[2].Name)
	}
}

func TestSortChildren_Recursive(t *testing.T) {
	root := NewDirectoryNode("root", "/root")
	
	child := NewDirectoryNode("child", "/root/child")
	grandchild1 := NewDirectoryNode("gc1", "/root/child/gc1")
	grandchild1.LOC = 5
	grandchild2 := NewDirectoryNode("gc2", "/root/child/gc2")
	grandchild2.LOC = 15
	
	child.AddChild(grandchild1)
	child.AddChild(grandchild2)
	root.AddChild(child)
	
	root.SortChildrenRecursive()
	
	// Grandchildren should also be sorted
	if child.Children[0].Name != "gc2" {
		t.Errorf("Expected first grandchild to be 'gc2' (LOC=15), got '%s'", child.Children[0].Name)
	}
	if child.Children[1].Name != "gc1" {
		t.Errorf("Expected second grandchild to be 'gc1' (LOC=5), got '%s'", child.Children[1].Name)
	}
}