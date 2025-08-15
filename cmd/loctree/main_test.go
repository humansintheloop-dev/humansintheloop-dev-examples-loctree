package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestApplicationBuilds(t *testing.T) {
	cmd := exec.Command("go", "build", "-o", "test_loctree", "main.go")
	cmd.Dir = "."
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build application: %v", err)
	}
	
	// Clean up
	os.Remove("test_loctree")
}

func TestMainIntegration_NoArguments(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	
	err := cmd.Run()
	if err == nil {
		t.Error("Expected error when no arguments provided")
	}
	
	output := stderr.String()
	if !strings.Contains(output, "Usage: loctree") {
		t.Errorf("Expected usage message, got: %s", output)
	}
}

func TestMainIntegration_ValidDirectory(t *testing.T) {
	// Skip TUI tests in CI environment
	// TUI requires interactive terminal which isn't available in test env
	t.Skip("Skipping TUI integration test")
}