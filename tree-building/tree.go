package tree

// Record holds the ID of the current record and that of its nearest parent
type Record struct {
	ID     int
	Parent int
}

// Node holds the id of the current node and a reference array its nearest descendants
type Node struct {
	ID       int
	Children []*Node
}

// Build creates node tree for an array of records
func Build(rs []Record) (*Node, error) {
	if len(rs) == 0 {
		return nil, nil
	}

	return nil, nil
}
