package fiveetools

import (
	"encoding/json"
	"fmt"

	"github.com/custompro98/r20conv/internal/pkg/adventure"
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
func Parse(b []byte) (adventure.Adventure, error) {
	var res adventure.Adventure
	var adv Adventure

	err := json.Unmarshal(b, &adv)

	if err != nil {
		return res, err
	}

	normalize(adv, &res)

	return res, nil
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
		var str string

		if err := json.Unmarshal(data, &str); err != nil {
			// TODO: this should more intelligently check for the right error type
		}

		se.Type = "text"
		se.Value = str
	} else {
		se.Type = res.Type
		se.Id = res.Id
		se.Entries = res.Entries
	}

	return nil
}

// normalize converts a local Adventure into the uniform Adventure type
func normalize(a Adventure, t *adventure.Adventure) {
	for _, v := range a.Sections {
		normalizeSection(v, &t.Sections)
	}
}

// normalizeSection converts a local Section into the uniform Section type
func normalizeSection(s Section, t *([]adventure.Section)) {
	sec := adventure.Section{
		Type: s.Type,
		Name: s.Name,
		Page: s.Page,
		Id:   s.Id,
	}

	for _, v := range s.Entries {
		normalizeEntry(v, &sec.Entries)
	}

	*t = append(*t, sec)
}

// normalizeEntry converts a local Entry into the uniform Entry type
// this function is called recursively to go down the nested structure of entries
func normalizeEntry(e Entry, t *([]adventure.Entry)) {
	fmt.Printf("normalizing entry: id: %v ; value: %v", e.Id, e.Value)
	ent := adventure.Entry{
		Type:  e.Type,
		Id:    e.Id,
		Value: e.Value,
	}

	for _, v := range e.Entries {
		normalizeEntry(v, &ent.Entries)
	}

	*t = append(*t, ent)
}

// isEmpty returns true if the input is empty
func isEmpty(s string) bool {
	return s == ""
}
