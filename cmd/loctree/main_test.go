package main

import (
	"os"
	"os/exec"
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