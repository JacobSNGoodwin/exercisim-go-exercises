package tree

// Record gives the id of a record and of it's direct parent
type Record struct {
	ID     int
	Parent int
}

// Build creates a tree of records based on parents and children
func Build(records []Record) (*Node, error) {

}
