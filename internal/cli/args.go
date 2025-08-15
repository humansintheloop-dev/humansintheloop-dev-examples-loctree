package cli

import (
	"fmt"
	"os"
)

// ParseArgs parses command-line arguments and returns the directory path
func ParseArgs(args []string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("Usage: loctree <directory_path>")
	}
	
	if len(args) != 1 {
		return "", fmt.Errorf("Expected exactly one argument, got %d", len(args))
	}
	
	return args[0], nil
}

// ValidatePath checks if the given path exists and is a directory
func ValidatePath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("Error: Path does not exist: %s", path)
		}
		return fmt.Errorf("Error accessing path: %v", err)
	}
	
	if !info.IsDir() {
		return fmt.Errorf("Error: Path is not a directory: %s", path)
	}
	
	return nil
}