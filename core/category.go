package core

// Category of item
type Category int

const (
	Lifetime Category = 1 << iota // 1 << 0 ==> 1
	Backup                        // 1 << 1 ==> 2
	Film                          // 1 << 2 ==> 4
	Music                         // 1 << 3 ==> 8
	Photo                         // 1 << 4 ==> 16

	AllCategories = Backup | Film | Music | Photo
)
