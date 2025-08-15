package scanner

import (
	"bufio"
	"os"
)

// CountLines counts the number of lines in a file
// Returns 0 for binary files
func CountLines(filePath string) (int, error) {
	// Check if file is binary first
	isBin, err := IsBinary(filePath)
	if err != nil {
		return 0, err
	}
	if isBin {
		return 0, nil
	}
	
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	
	return lines, nil
}