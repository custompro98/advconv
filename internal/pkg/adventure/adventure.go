package adventure

import (
	"encoding/json"
)

// Adventure encapsulates the entire adventure object from 5eTools
type Adventure struct {
	Sections []Section `json:"data"`
}

// Section encapsulates each arbitrary section in the Adventure
type Section struct {
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Page    int     `json:"page"`
	Id      string  `json:"id"`
	Entries []Entry `json:"entries"`
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

// EntryJSON is a helper struct to parse dynamic Entry formats:
// A) Edge Branch: { Type, Id, []Entries }
// B) Leaf Node: String
type EntryJSON struct {
	Type    string  `json:"type"`
	Id      string  `json:"id"`
	Entries []Entry `json:"entries"`
}

// Parse turns a JSON string from 5eTools into an @Adventure
func Parse(s string) (Adventure, error) {
	var adventure Adventure

	err := json.Unmarshal([]byte(s), &adventure)

	if err != nil {
		return Adventure{}, err
	}

	return adventure, nil
}

// UnmarshalJSON is a custom JSON parser to handle the edge branch
// and leaf node structure of the @Entry // @EntryJSON struct
func (se *Entry) UnmarshalJSON(data []byte) error {
	var res EntryJSON

	if err := json.Unmarshal(data, &res); err != nil {
		// TODO: this should more intelligently check for the right error type
	}

	// We found a leaf node
	if isEmpty(res.Type) {
		se.Type = "text"
		se.Value = string(data)
	} else {
		se.Type = res.Type
		se.Id = res.Id
		se.Entries = res.Entries
	}

	return nil
}

// isEmpty returns true if the input is empty
func isEmpty(s string) bool {
	return s == ""
}
