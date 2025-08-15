package tree

import "sort"

// DirectoryNode represents a directory in the tree structure
type DirectoryNode struct {
	Name       string
	Path       string
	LOC        int              // Total LOC (including children)
	FileLOC    int              // LOC from files in this directory only
	Children   []*DirectoryNode
	IsExpanded bool
	Parent     *DirectoryNode
}

// NewDirectoryNode creates a new directory node
func NewDirectoryNode(name, path string) *DirectoryNode {
	return &DirectoryNode{
		Name:       name,
		Path:       path,
		LOC:        0,
		FileLOC:    0,
		Children:   []*DirectoryNode{},
		IsExpanded: false,
		Parent:     nil,
	}
}

// AddChild adds a child node and establishes parent-child relationship
func (n *DirectoryNode) AddChild(child *DirectoryNode) {
	n.Children = append(n.Children, child)
	child.Parent = n
}

// CalculateLOC recursively calculates the total LOC for this node and all children
func (n *DirectoryNode) CalculateLOC() {
	// Start with files in this directory
	n.LOC = n.FileLOC
	
	// Recursively calculate for children and add to total
	for _, child := range n.Children {
		child.CalculateLOC()
		n.LOC += child.LOC
	}
}

// SortChildren sorts the immediate children by LOC (descending)
func (n *DirectoryNode) SortChildren() {
	sort.Slice(n.Children, func(i, j int) bool {
		return n.Children[i].LOC > n.Children[j].LOC
	})
}

// SortChildrenRecursive sorts all children and their descendants by LOC (descending)
func (n *DirectoryNode) SortChildrenRecursive() {
	n.SortChildren()
	for _, child := range n.Children {
		child.SortChildrenRecursive()
	}
}