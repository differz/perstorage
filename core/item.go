package core

// Item contains single file
type Item struct {
	ID         int
	Name       string
	Size       int
	Available  bool
	Filename   string
	SourceName string
}

// IsNew returns true if it is new item
func (i Item) IsNew() bool {
	return i.ID == 0
}
