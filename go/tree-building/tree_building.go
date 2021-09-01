package tree

// Record type for tree input
type Record struct {
	ID     int
	Parent int
}

// Node type for tree output
type Node struct {
	ID       int
	Children []*Node
}

// Build records function
func Build(records []Record) (*Node, error) {
	// Nodes have an ID (the parent ID) and then a slice of pointers to nodes
	// We need to go through the slice of records and make a node which takes the parent ID as an int key and the ID as a pointer which is put into a slice
	var n *Node
	var c []*Node

	for _, r := range records {
		n.ID = r.Parent
		c[0].ID = r.ID
		c[0].Children = nil
		n.Children = c

	}
	return n, nil
}
