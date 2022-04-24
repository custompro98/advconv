package fiveetools

import (
	"github.com/custompro98/advconv/internal/pkg/adventure"
)

// denormalize converts an @adventure.Adventure into the local @Adventure type
func denormalize(a adventure.Adventure, t *Adventure) {
	t.Sections = make([]Section, 0)

	// need to do the opposite of normalizing
	for _, v := range a.Sections {
		denormalizeSection(v, &t.Sections)
	}
}

// denormalizeSection converts an @adventure.Section into the local @Section type
func denormalizeSection(s adventure.Section, t *([]Section)) {
	sec := Section{
		Type:    s.Type,
		Name:    s.Name,
		Page:    s.Page,
		Id:      s.Id,
		Entries: make([]Entry, 0),
	}

	for _, v := range s.Entries {
		denormalizeEntry(v, &sec.Entries)
	}

	*t = append(*t, sec)
}

// denormalizeEntry converts an @adventure.Entry into the local @Entry type
// this function is called recursively to go down the nested structure of entries
func denormalizeEntry(e adventure.Entry, t *([]Entry)) {
	ent := Entry{
		Type:  e.Type,
		Id:    e.Id,
		Value: e.Value,
	}

	for _, v := range e.Entries {
		denormalizeEntry(v, &ent.Entries)
	}

	*t = append(*t, ent)
}
