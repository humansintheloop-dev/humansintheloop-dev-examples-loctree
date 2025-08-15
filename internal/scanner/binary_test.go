package scanner

import (
	"path/filepath"
	"testing"
)

func TestIsBinary_TextFile(t *testing.T) {
	textFile := filepath.Join("testdata", "text_file.txt")
	result, err := IsBinary(textFile)
	if err != nil {
		t.Fatalf("Error checking binary status: %v", err)
	}
	if result {
		t.Error("Expected text file to be detected as non-binary")
	}
}

func TestIsBinary_BinaryFile(t *testing.T) {
	binaryFile := filepath.Join("testdata", "binary_file.bin")
	result, err := IsBinary(binaryFile)
	if err != nil {
		t.Fatalf("Error checking binary status: %v", err)
	}
	if !result {
		t.Error("Expected binary file to be detected as binary")
	}
}

func TestIsBinary_EmptyFile(t *testing.T) {
	emptyFile := filepath.Join("testdata", "empty_file.txt")
	result, err := IsBinary(emptyFile)
	if err != nil {
		t.Fatalf("Error checking binary status: %v", err)
	}
	if result {
		t.Error("Expected empty file to be detected as non-binary")
	}
}

func TestIsBinary_NonExistentFile(t *testing.T) {
	nonExistent := filepath.Join("testdata", "does_not_exist.txt")
	_, err := IsBinary(nonExistent)
	if err == nil {
		t.Error("Expected error for non-existent file")
	}
}