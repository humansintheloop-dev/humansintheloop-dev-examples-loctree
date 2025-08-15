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
	cmd := exec.Command("go", "run", "main.go", "/tmp")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	
	err := cmd.Run()
	if err != nil {
		t.Errorf("Expected no error for valid directory, got: %v", err)
	}
	
	output := stdout.String()
	if !strings.Contains(output, "Scanning: /tmp") {
		t.Errorf("Expected 'Scanning: /tmp' in output, got: %s", output)
	}
	if !strings.Contains(output, "Total LOC:") {
		t.Errorf("Expected 'Total LOC:' in output, got: %s", output)
	}
	if !strings.Contains(output, "Tree structure:") {
		t.Errorf("Expected 'Tree structure:' in output, got: %s", output)
	}
}