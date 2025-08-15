package tree

import (
	"os"
	"path/filepath"
	"strings"
	
	"github.com/user/loctree/internal/scanner"
)

// BuildTree builds a directory tree with LOC information
func BuildTree(rootPath string) (*DirectoryNode, error) {
	// Verify path exists
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return nil, os.ErrNotExist
	}
	
	// Create root node
	rootName := filepath.Base(rootPath)
	root := NewDirectoryNode(rootName, rootPath)
	
	// Map to store nodes by path for quick lookup
	nodeMap := make(map[string]*DirectoryNode)
	nodeMap[rootPath] = root
	
	// Walk directory tree
	err = filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // Skip errors
		}
		
		// Skip root (already created)
		if path == rootPath {
			return nil
		}
		
		// Skip hidden directories
		if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
			return filepath.SkipDir
		}
		
		// Get parent path
		parentPath := filepath.Dir(path)
		parentNode, exists := nodeMap[parentPath]
		if !exists {
			// Parent doesn't exist (shouldn't happen in normal walk)
			return nil
		}
		
		if d.IsDir() {
			// Create directory node
			node := NewDirectoryNode(d.Name(), path)
			parentNode.AddChild(node)
			nodeMap[path] = node
		} else {
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
			
			// Count lines in file and add to parent's FileLOC
			lines, err := scanner.CountLines(path)
			if err != nil {
				return nil // Skip files we can't read
			}
			parentNode.FileLOC += lines
		}
		
		return nil
	})
	
	if err != nil {
		return nil, err
	}
	
	// Calculate total LOC for all nodes
	root.CalculateLOC()
	
	return root, nil
}