// Package tree implements methods to create a tree of nodes.
package tree

import (
	"errors"
	"sort"
)

// Record reprsents a leaf of a node.
type Record struct {
	ID     int
	Parent int
}

// Node represents a node of a tree.
type Node struct {
	ID       int
	Children []*Node
}

// Build creates a tree represented as a list of nodes and returns a pointer to the head of the list.
func Build(records []Record) (*Node, error) {
	nodes := map[int]*Node{}
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		if r.ID != i || r.Parent > r.ID || r.ID > 0 && r.Parent == r.ID {
			return nil, errors.New("invalid tree")
		}
		nodes[i] = &Node{ID: r.ID}
		if r.ID > 0 {
			parent := nodes[r.Parent]
			parent.Children = append(parent.Children, nodes[i])
		}

	}
	return nodes[0], nil
}
