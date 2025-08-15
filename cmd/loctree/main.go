package main

import (
	"fmt"
	"os"
	"strings"
	
	"github.com/user/loctree/internal/cli"
	"github.com/user/loctree/internal/tree"
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
	
	// Build the tree
	fmt.Printf("Scanning: %s\n", path)
	root, err := tree.BuildTree(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error building tree: %v\n", err)
		os.Exit(1)
	}
	
	// Display tree structure (text format for now)
	fmt.Printf("\nTree structure:\n")
	printTree(root, 0)
	fmt.Printf("\nTotal LOC: %d\n", root.LOC)
}

// printTree recursively prints the tree structure
func printTree(node *tree.DirectoryNode, depth int) {
	indent := strings.Repeat("  ", depth)
	indicator := ""
	if len(node.Children) > 0 {
		if node.IsExpanded {
			indicator = "▼ "
		} else {
			indicator = "▶ "
		}
	}
	fmt.Printf("%s%s%d %s\n", indent, indicator, node.LOC, node.Name)
	
	// For text display, show all children (simulate expanded)
	for _, child := range node.Children {
		printTree(child, depth+1)
	}
}