package tree

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	// Nodes have an ID (the parent ID) and then a slice of pointers to nodes
	// We need to go through the slice of records and make a node which takes the parent ID as an int key and the ID as a pointer which is put into a slice
	var n Node

	for _, r := range records {
		n.ID = r.Parent
		childID := []&r.ID
		n.Children = append(n.Children, childID)
	}
	return n, nil
}
