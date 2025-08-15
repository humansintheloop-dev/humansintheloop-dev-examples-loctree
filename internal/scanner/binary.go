package scanner

import (
	"os"
)

// IsBinary checks if a file is binary by looking for null bytes
func IsBinary(filePath string) (bool, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	
	// Read first 512 bytes
	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return false, nil // Treat read errors as non-binary
	}
	
	// Check for null bytes
	for i := 0; i < n; i++ {
		if buffer[i] == 0 {
			return true, nil
		}
	}
	
	return false, nil
}