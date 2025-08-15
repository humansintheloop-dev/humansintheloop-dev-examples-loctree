package tree

// ToggleExpanded toggles the expanded state of a directory node
func (n *DirectoryNode) ToggleExpanded() {
	n.IsExpanded = !n.IsExpanded
}

// GetVisibleNodes returns a flat list of currently visible nodes
func GetVisibleNodes(root *DirectoryNode) []*DirectoryNode {
	var visible []*DirectoryNode
	addVisibleNodes(root, &visible)
	return visible
}

// addVisibleNodes recursively adds visible nodes to the list
func addVisibleNodes(node *DirectoryNode, visible *[]*DirectoryNode) {
	*visible = append(*visible, node)
	
	if node.IsExpanded {
		for _, child := range node.Children {
			addVisibleNodes(child, visible)
		}
	}
}