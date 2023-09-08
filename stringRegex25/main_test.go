package main

import (
	"testing"

	"gotest.tools/assert"
)

func TestGetMatchedNames(t *testing.T) {

	testData := []struct {
		description string
		names       []string
		expected    []string
		matchTo     string
	}{
		{"Valid regex", []string{"rAnil", "Aniket", "Pratik", "12345"}, []string{"rAnil", "Aniket"}, "ni"},
		{"Invalid regex", []string{"Merce", "remiges", "12345"}, []string{"12345"}, "23"},
		{"Empty slice", []string{}, []string(nil), "varun"},
	}

	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			actual := getMatchedNames(test.matchTo, test.names)
			assert.DeepEqual(t, test.expected, actual)
		})
	}
}
