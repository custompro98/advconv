package fiveetools

import (
	"encoding/json"

	"github.com/custompro98/r20conv/internal/pkg/adventure"
)

// Adventure encapsulates the entire adventure object from 5eTools
type Adventure struct {
	Sections []Section `json:"data"`
}

// Section encapsulates each arbitrary section in the Adventure
type Section struct {
	Type    string  `json:"type,omitempty"`
	Name    string  `json:"name,omitempty"`
	Page    int     `json:"page,omitempty"`
	Id      string  `json:"id,omitempty"`
	Entries []Entry `json:"entries,omitempty"`
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
	Type    string  `json:"type,omitempty"`
	Id      string  `json:"id,omitempty"`
	Entries []Entry `json:"entries,omitempty"`
}

// Parse turns a JSON string from 5eTools into an @adventure.Adventure
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

// Serialize turns an @adventure.Adventure into a []byte for writing to a file
func Serialize(a adventure.Adventure) ([]byte, error) {
	var res []byte
	var adv Adventure

	// denormalize into a local Adventure
	denormalize(a, &adv)

	// Marshal the local Adventure to []byte
	res, err := json.Marshal(adv)

	if err != nil {
		// TODO: this should more intelligently check for the right error type
	}

	return res, nil
}

// MarshalJSON is a custom JSON parser to handle the edge branch
// and leaf node structure of the @Entry // @EntryJSON struct
func (se *Entry) MarshalJSON() ([]byte, error) {
	if se.Type == "text" {
		return json.Marshal(se.Value)
	}

	entryJSON := EntryJSON{
		Type:    se.Type,
		Id:      se.Id,
		Entries: se.Entries,
	}

	return json.Marshal(entryJSON)
}

// isEmpty returns true if the input is empty
func isEmpty(s string) bool {
	return s == ""
}
