package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

// ScanResult holds the results of scanning a directory
type ScanResult struct {
	TotalLOC     int
	FilesScanned int
	DirsScanned  int
}

// ScanDirectory recursively scans a directory and counts lines of code
func ScanDirectory(dirPath string) (*ScanResult, error) {
	// Check if directory exists
	info, err := os.Stat(dirPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, os.ErrNotExist
	}
	
	result := &ScanResult{}
	
	err = filepath.WalkDir(dirPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			// Skip directories we can't read
			return nil
		}
		
		// Skip hidden directories
		if d.IsDir() {
			result.DirsScanned++
			if strings.HasPrefix(d.Name(), ".") && d.Name() != "." {
				return filepath.SkipDir
			}
			return nil
		}
		
		// Skip symbolic links
		info, err := d.Info()
		if err != nil {
			return nil
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}
		
		// Skip hidden files
		if strings.HasPrefix(d.Name(), ".") {
			return nil
		}
		
		// Count lines in regular files
		lines, err := CountLines(path)
		if err != nil {
			// Skip files we can't read
			return nil
		}
		
		result.FilesScanned++
		result.TotalLOC += lines
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	return result, nil
}