package main

import (
	"fmt"
	"os"
	
	"github.com/user/loctree/internal/cli"
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
	
	// For now, just print the path we're scanning
	fmt.Printf("Scanning: %s\n", path)
}