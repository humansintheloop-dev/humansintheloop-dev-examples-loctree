package scanner

import (
	"path/filepath"
	"testing"
)

func TestScanDirectory_SingleFile(t *testing.T) {
	// Test with just the single_line.txt file's parent directory
	singleFile := filepath.Join("testdata")
	result, err := ScanDirectory(singleFile)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}
	// We have: text_file.txt (5), single_line.txt (1), empty_file.txt (0), binary_file.bin (0)
	// Plus test_project with main.go (5) and src/utils.go (5)
	// Total = 5 + 1 + 0 + 0 + 5 + 5 = 16
	if result.TotalLOC != 16 {
		t.Errorf("Expected 16 total lines, got %d", result.TotalLOC)
	}
}

func TestScanDirectory_MultipleFiles(t *testing.T) {
	testProject := filepath.Join("testdata", "test_project")
	result, err := ScanDirectory(testProject)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}
	// main.go has 5 lines, src/utils.go has 5 lines
	// binary.bin should be 0, .hidden should be ignored
	if result.TotalLOC != 10 {
		t.Errorf("Expected 10 total lines, got %d", result.TotalLOC)
	}
}

func TestScanDirectory_MixedContent(t *testing.T) {
	testProject := filepath.Join("testdata", "test_project")
	result, err := ScanDirectory(testProject)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}
	// Should handle both text and binary files
	if result.TotalLOC != 10 {
		t.Errorf("Expected 10 total lines (text files only), got %d", result.TotalLOC)
	}
}

func TestScanDirectory_HiddenFiles(t *testing.T) {
	// Hidden directories should be ignored
	testProject := filepath.Join("testdata", "test_project")
	result, err := ScanDirectory(testProject)
	if err != nil {
		t.Fatalf("Error scanning directory: %v", err)
	}
	// .hidden directory should be ignored
	if result.FilesScanned > 3 {
		t.Errorf("Expected at most 3 files scanned (excluding hidden), got %d", result.FilesScanned)
	}
}

func TestScanDirectory_NonExistent(t *testing.T) {
	nonExistent := filepath.Join("testdata", "does_not_exist")
	_, err := ScanDirectory(nonExistent)
	if err == nil {
		t.Error("Expected error for non-existent directory")
	}
}