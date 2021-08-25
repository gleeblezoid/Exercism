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
	// We need to go through the slice of records and get the parent ID for each []record
	// If the parent is one we already have then it should be the key to a map of node pointers
	// The value for that map is the slice of child node pointers which need to be constructed based on the parent ID
}
