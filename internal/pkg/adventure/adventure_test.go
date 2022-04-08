package adventure

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		scenario    string
		input       string
		expected    Adventure
		errExpected bool
	}{
		{
			scenario: "An empty input produces an empty adventure",
			input:    "{}",
			expected: Adventure{},
		},
		{
			scenario: "The opening data tag is ignored",
			input:    "{\"data\": []}",
			expected: Adventure{},
		},
		{
			scenario: "The adventure can be broken up into many named sections",
			input:    "{\"data\": [{\"type\": \"section\", \"name\": \"Section 1\"},{\"type\": \"section\", \"name\": \"Section 2\"}]}",
			expected: Adventure{
				Sections: []Section{
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
			expected: Adventure{
				Sections: []Section{
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
			expected: Adventure{
				Sections: []Section{
					{
						Entries: []Entry{
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
			expected: Adventure{
				Sections: []Section{
					{
						Entries: []Entry{
							{
								Type: "entry",
								Entries: []Entry{
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
			result, err := Parse(v.input)

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
			}
		})
	}
}
