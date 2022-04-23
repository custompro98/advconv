package adventure

// Adventure encapsulates the entire adventure object from 5eTools
type Adventure struct {
	Sections []Section
}

// Section encapsulates each arbitrary section in the Adventure
type Section struct {
	Type    string
	Name    string
	Page    int
	Id      string
	Entries []Entry
}

// Section Entry encapsulates each entry in a Section
type Entry struct {
	Type string
	Id   string
	// Value is only present if Entries is not, it's a leaf node in the tree
	Value string
	// Entries is only present if Value is not, it's a branch in the tree
	Entries []Entry
}
