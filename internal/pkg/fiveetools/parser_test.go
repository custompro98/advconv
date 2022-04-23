package fiveetools

import (
	"testing"

	"github.com/custompro98/r20conv/internal/pkg/adventure"
)

func TestParse(t *testing.T) {
	tests := []struct {
		scenario    string
		input       string
		expected    adventure.Adventure
		errExpected bool
	}{
		{
			scenario: "An empty input produces an empty adventure",
			input:    "{}",
			expected: adventure.Adventure{},
		},
		{
			scenario: "The opening data tag is ignored",
			input:    "{\"data\": []}",
			expected: adventure.Adventure{},
		},
		{
			scenario: "The adventure can be broken up into many named sections",
			input:    "{\"data\": [{\"type\": \"section\", \"name\": \"Section 1\"},{\"type\": \"section\", \"name\": \"Section 2\"}]}",
			expected: adventure.Adventure{
				Sections: []adventure.Section{
					{
						Type: "section",
						Name: "Section 1",
					},
					{
						Type: "section",
						Name: "Section 2",
					},
				},
			},
		},
		{
			scenario: "The named sections have types, page numbers, and ids",
			input:    "{\"data\": [{\"type\": \"section\", \"name\": \"Section 1\", \"page\": 1, \"id\": \"abc123\"}]}",
			expected: adventure.Adventure{
				Sections: []adventure.Section{
					{
						Type: "section",
						Name: "Section 1",
						Page: 1,
						Id:   "abc123",
					},
				},
			},
		},
		{
			scenario: "The named sections also have entries which can be strings",
			input:    "{\"data\": [{\"entries\": [\"Some description\"]}]}",
			expected: adventure.Adventure{
				Sections: []adventure.Section{
					{
						Entries: []adventure.Entry{
							{
								Type:  "text",
								Value: "Some description",
							},
						},
					},
				},
			},
		},
		{
			scenario: "The named sections can also have entries which are full objects",
			input:    "{\"data\": [{\"entries\": [{\"type\": \"entry\", \"entries\": [\"Some description\"]}]}]}",
			expected: adventure.Adventure{
				Sections: []adventure.Section{
					{
						Entries: []adventure.Entry{
							{
								Type: "entry",
								Entries: []adventure.Entry{
									{
										Type:  "text",
										Value: "Some description",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, v := range tests {
		t.Run(v.scenario, func(t *testing.T) {
			result, err := Parse([]byte(v.input))

			if err != nil && !v.errExpected {
				t.Error(err)
			}

			if len(result.Sections) != len(v.expected.Sections) {
				t.Errorf("len(Sections) expected %v, got %v", v.expected, result)
			}

			for i, section := range result.Sections {
				expected := v.expected.Sections[i]

				if section.Type != expected.Type {
					t.Errorf("section.Type expected %v, got %v", expected.Type, section.Type)
				}

				if section.Name != expected.Name {
					t.Errorf("section.Name expected %v, got %v", expected.Name, section.Name)
				}

				if section.Page != expected.Page {
					t.Errorf("section.Page expected %v, got %v", expected.Page, section.Page)
				}

				if section.Id != expected.Id {
					t.Errorf("section.Id expected %v, got %v", expected.Id, section.Id)
				}

				if len(section.Entries) != len(expected.Entries) {
					t.Errorf("len(Entries) expected %v, got %v", expected.Entries, section.Entries)
				}

				for j, entry := range section.Entries {
					expected := v.expected.Sections[i].Entries[j]

					if entry.Id != expected.Id {
						t.Errorf("entry.Id expected %v, got %v", expected.Id, entry.Id)
					}

					if entry.Value != expected.Value {
						t.Errorf("entry.Value expected %v, got %v", expected.Value, entry.Value)
					}

					if entry.Value != expected.Value {
						t.Errorf("entry.Value expected %v, got %v", expected.Value, entry.Value)
					}

					if len(entry.Entries) != len(expected.Entries) {
						t.Errorf("len(Entries) expected %v, got %v", expected.Entries, section.Entries)
					}

					for k, subEntry := range entry.Entries {
						expected := v.expected.Sections[i].Entries[j].Entries[k]

						if subEntry.Id != expected.Id {
							t.Errorf("subEntry.Id expected %v, got %v", expected.Id, subEntry.Id)
						}

						if subEntry.Value != expected.Value {
							t.Errorf("subEntry.Value expected %v, got %v", expected.Value, subEntry.Value)
						}

						if subEntry.Value != expected.Value {
							t.Errorf("subEntry.Value expected %v, got %v", expected.Value, subEntry.Value)
						}
					}
				}
			}
		})
	}
}

/* func Test_Serialize(t *testing.T) {
	tests := []struct {
		scenario    string
		input       adventure.Adventure
		expected    string
		errExpected bool
	}{
		{
			scenario: "An empty adventure creates an empty string",
			input: adventure.Adventure{
				Sections: []adventure.Section{},
			},
			expected: "{\"data\": []}",
		},
	}

	for _, v := range tests {
		t.Run(v.scenario, func(t *testing.T) {
			result, err := Serialize(v.input)

			if err != nil && !v.errExpected {
				t.Error(err)
			}

			if result != strings.ReplaceAll(v.expected, " ", "") {
				t.Errorf("result expected %v, got %v", v.expected, result)
			}
		})
	}

} */
