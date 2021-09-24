package tree

import {
	"sort"
	"errors"
}
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

func Build(records []Record) (*Node, error) {
	// If there's no input at all then error out
	if len(records) == 0 {
		return nil, nil
	}
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	if records[0].Parent != 0 || records[0].ID != 0 {
		return nil, errors.New("no root node")
	}
	nodes := make([]Node, len(records))
	for i, record := range records[1:] {
		if i+1 != record.ID {
			return nil, errors.New("duplicate or non-continuous node")
		}
		if i+1 <= record.Parent {
			return nil, errors.New("direct/indirect cycle in the tree")
		}
		nodes[i+1].ID = i + 1
		children := nodes[record.Parent].Children
		children = append(children, &nodes[i+1])
		nodes[record.Parent].Children = children
	}
	return &nodes[0], nil
}
