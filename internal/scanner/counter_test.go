package scanner

import (
	"path/filepath"
	"testing"
)

func TestCountLines_EmptyFile(t *testing.T) {
	emptyFile := filepath.Join("testdata", "empty_file.txt")
	count, err := CountLines(emptyFile)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}
	if count != 0 {
		t.Errorf("Expected 0 lines for empty file, got %d", count)
	}
}

func TestCountLines_SingleLine(t *testing.T) {
	singleFile := filepath.Join("testdata", "single_line.txt")
	count, err := CountLines(singleFile)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 line for single line file, got %d", count)
	}
}

func TestCountLines_MultipleLines(t *testing.T) {
	textFile := filepath.Join("testdata", "text_file.txt")
	count, err := CountLines(textFile)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}
	if count != 5 {
		t.Errorf("Expected 5 lines for text file, got %d", count)
	}
}

func TestCountLines_BinaryFile(t *testing.T) {
	binaryFile := filepath.Join("testdata", "binary_file.bin")
	count, err := CountLines(binaryFile)
	if err != nil {
		t.Fatalf("Error counting lines: %v", err)
	}
	if count != 0 {
		t.Errorf("Expected 0 lines for binary file, got %d", count)
	}
}

func TestCountLines_NonExistentFile(t *testing.T) {
	nonExistent := filepath.Join("testdata", "does_not_exist.txt")
	count, err := CountLines(nonExistent)
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
	if count != 0 {
		t.Errorf("Expected 0 lines for non-existent file, got %d", count)
	}
}