package tree

import (
	"path/filepath"
	"testing"
)

func TestBuildTree_SingleDirectory(t *testing.T) {
	// Use the test data from scanner package
	testPath := filepath.Join("..", "scanner", "testdata", "test_project")
	
	tree, err := BuildTree(testPath)
	if err != nil {
		t.Fatalf("Error building tree: %v", err)
	}
	
	if tree.Name != "test_project" {
		t.Errorf("Expected root name 'test_project', got '%s'", tree.Name)
	}
	
	if tree.Path != testPath {
		t.Errorf("Expected root path '%s', got '%s'", testPath, tree.Path)
	}
}

func TestBuildTree_NestedDirectories(t *testing.T) {
	testPath := filepath.Join("..", "scanner", "testdata", "test_project")
	
	tree, err := BuildTree(testPath)
	if err != nil {
		t.Fatalf("Error building tree: %v", err)
	}
	
	// Should have src directory as child
	srcFound := false
	for _, child := range tree.Children {
		if child.Name == "src" {
			srcFound = true
			// src directory should have LOC from its files
			if child.FileLOC == 0 {
				t.Error("Expected src directory to have FileLOC > 0")
			}
			break
		}
	}
	
	if !srcFound {
		t.Error("Expected to find 'src' directory in children")
	}
}

func TestBuildTree_CalculatesLOC(t *testing.T) {
	testPath := filepath.Join("..", "scanner", "testdata", "test_project")
	
	tree, err := BuildTree(testPath)
	if err != nil {
		t.Fatalf("Error building tree: %v", err)
	}
	
	// Total LOC should be calculated
	if tree.LOC == 0 {
		t.Error("Expected non-zero LOC for tree")
	}
	
	// Should match what scanner found (10 lines total)
	// main.go has 5 lines, src/utils.go has 5 lines
	if tree.LOC != 10 {
		t.Errorf("Expected total LOC of 10, got %d", tree.LOC)
	}
}

func TestBuildTree_IgnoresHiddenDirectories(t *testing.T) {
	testPath := filepath.Join("..", "scanner", "testdata", "test_project")
	
	tree, err := BuildTree(testPath)
	if err != nil {
		t.Fatalf("Error building tree: %v", err)
	}
	
	// Should not have .hidden directory as child
	for _, child := range tree.Children {
		if child.Name == ".hidden" {
			t.Error("Hidden directory should not be in tree")
		}
	}
}

func TestBuildTree_NonExistentPath(t *testing.T) {
	_, err := BuildTree("/path/does/not/exist")
	if err == nil {
		t.Error("Expected error for non-existent path")
	}
}