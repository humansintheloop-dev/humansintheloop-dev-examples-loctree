package main

import (
	"fmt"
	"os"
	
	"github.com/user/loctree/internal/cli"
	"github.com/user/loctree/internal/scanner"
)

func main() {
	// Parse command-line arguments
	path, err := cli.ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	
	// Validate the path
	err = cli.ValidatePath(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	
	// Scan the directory
	fmt.Printf("Scanning: %s\n", path)
	result, err := scanner.ScanDirectory(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning directory: %v\n", err)
		os.Exit(1)
	}
	
	// Display results
	fmt.Printf("Total LOC: %d\n", result.TotalLOC)
	fmt.Printf("Files scanned: %d\n", result.FilesScanned)
	fmt.Printf("Directories scanned: %d\n", result.DirsScanned)
}