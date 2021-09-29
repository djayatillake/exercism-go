package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func validate_records(records []Record) error {
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		switch {
		case i == 0 && r.ID != 0:
			return errors.New("no root node")
		case r.ID == 0 && r.Parent != 0:
			return errors.New("root node has parent")
		case i > 0 && r.ID == 0:
			return errors.New("duplicate root")
		case i > 0 && r.ID < i:
			return errors.New("duplicate node")
		case i > 0 && r.ID > i:
			return errors.New("non-continuous")
		case r.ID <= r.Parent && i > 0:
			return errors.New("cycle or higher id parent of lower id")
		}
	}
	return nil
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	err := validate_records(records)
	if err != nil {
		return nil, err
	}
	nodes := map[int]*Node{}
	for _, r := range records {
		nodes[r.ID] = &Node{ID: r.ID}

		if r.ID == 0 {
			continue
		}

		nodes[r.Parent].Children = append(nodes[r.Parent].Children, nodes[r.ID])
	}
	return nodes[0], nil
}
