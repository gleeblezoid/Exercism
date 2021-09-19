package tree

import "sort"

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
	if len(records) == 0 {
		return nil, nil
	}
	idList = orderedList(records)
	n = makeChildren(idList, records)
	return n, nil
}

// Grab the parent ints from the records and turn into an ordered list
func orderedList(records []Record) []int {
	idList := make([]int)
	for _, r := range records {
		idList = append(idList, r.Parent)
	}
	idList = sort.IntSlice(idList)
	return idList
}

// For each parent ID make a node struct out of records
func makeChildren(idList []int, records []Record) []*Node {
	var nodes []*Node
	for _, v := range idList {
		var child Node
		child.ID = v
		for _, r := range records {
			if r.Parent == v {
				child.Children = append(child.Children, &r.ID)
			}
		}
		nodes = append(nodes, &child)
	}
	return nodes
}

// Nest the nodes by parent hierarchy
func nestNodes([]Node) []*Node {
	return nil
}
