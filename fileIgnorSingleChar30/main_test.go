package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestIsCommonWord(t *testing.T) {
	commonWords := []string{"the", "a", "an", "to", "of", "and", "in", "on", "is", "it"}

	testData := []struct {
		description string
		wordToLook  string
		expected    bool
	}{
		{"Known common word", "the", true},
		{"Unknown common word", "hello", false},
		{"Empty word", "", false},
	}

	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			actual := isCommonWord(test.wordToLook, &commonWords)
			assert.Equal(t, test.expected, actual)
		})
	}
}
