package cli

import (
	"fmt"
	"os"
)

func ParseArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("Usage: loctree <directory_path>")
	}
	
	if len(args) != 1 {
		return "", fmt.Errorf("Expected exactly one argument, got %d", len(args))
	}
	
	path := args[0]
	
	// Basic validation - check if path exists
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("Error: Path does not exist: %s", path)
		}
		return "", fmt.Errorf("Error accessing path: %v", err)
	}
	
	// Check if it's a directory
	if !info.IsDir() {
		return "", fmt.Errorf("Error: Path is not a directory: %s", path)
	}
	
	return path, nil
}