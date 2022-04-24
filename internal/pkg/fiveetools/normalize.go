package fiveetools

import (
	"github.com/custompro98/advconv/internal/pkg/adventure"
)

// normalize converts a local @Adventure into the uniform @adventure.Adventure type
func normalize(a Adventure, t *adventure.Adventure) {
	for _, v := range a.Sections {
		normalizeSection(v, &t.Sections)
	}
}

// normalizeSection converts a local @Section into the uniform @adventure.Section type
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

// normalizeEntry converts a local @Entry into the uniform @adventure.Entry type
// this function is called recursively to go down the nested structure of entries
func normalizeEntry(e Entry, t *([]adventure.Entry)) {
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
