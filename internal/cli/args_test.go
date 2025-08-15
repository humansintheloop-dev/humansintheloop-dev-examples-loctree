package cli

import (
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
		t.Errorf("Expected no error for valid directory, got: %v", err)
	}
	if path != "/tmp" {
		t.Errorf("Expected path to be '/tmp', got: %s", path)
	}
}

func TestParseArgs_InvalidPath(t *testing.T) {
	args := []string{"/this/path/does/not/exist/at/all"}
	_, err := ParseArgs(args)
	if err == nil {
		t.Error("Expected error for non-existent path, got nil")
	}
}