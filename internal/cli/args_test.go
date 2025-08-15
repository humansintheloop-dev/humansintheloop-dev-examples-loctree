package cli

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestParseArgs_NoArguments(t *testing.T) {
	args := []string{}
	_, err := ParseArgs(args)
	if err == nil {
		t.Error("Expected error when no arguments provided, got nil")
	}
}

func TestParseArgs_ValidDirectory(t *testing.T) {
	args := []string{"/tmp"}
	path, err := ParseArgs(args)
	if err != nil {
		t.Errorf("Expected no error for valid argument, got: %v", err)
	}
	if path != "/tmp" {
		t.Errorf("Expected path to be '/tmp', got: %s", path)
	}
}

func TestParseArgs_TooManyArguments(t *testing.T) {
	args := []string{"/tmp", "/usr"}
	_, err := ParseArgs(args)
	if err == nil {
		t.Error("Expected error for too many arguments, got nil")
	}
}

func TestValidatePath_DirectoryExists(t *testing.T) {
	// Create temp directory
	tempDir, err := ioutil.TempDir("", "loctree_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tempDir)
	
	err = ValidatePath(tempDir)
	if err != nil {
		t.Errorf("Expected no error for existing directory, got: %v", err)
	}
}

func TestValidatePath_FileNotDirectory(t *testing.T) {
	// Create temp file
	tempFile, err := ioutil.TempFile("", "loctree_test")
	if err != nil {
		t.Fatal(err)
	}
	tempFile.Close()
	defer os.Remove(tempFile.Name())
	
	err = ValidatePath(tempFile.Name())
	if err == nil {
		t.Error("Expected error for file (not directory), got nil")
	}
}

func TestValidatePath_NonExistentPath(t *testing.T) {
	nonExistentPath := filepath.Join(os.TempDir(), "this_does_not_exist_at_all_12345")
	err := ValidatePath(nonExistentPath)
	if err == nil {
		t.Error("Expected error for non-existent path, got nil")
	}
}