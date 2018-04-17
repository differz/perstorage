package core

// Item ...
type Item struct {
	ID         int
	Name       string
	Size       int
	Available  bool
	Filename   string
	SourceName string
}

// IsNew ...
func (i Item) IsNew() bool {
	return i.ID == 0
}
