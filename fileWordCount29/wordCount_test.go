package main

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestWordReader(t *testing.T) {

	testData := []struct {
		description string
		filePath    string
		expected    map[string]int
		err         string
	}{
		{"File with actual data", "/media/merceadm/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Downloads/input.txt", map[string]int{"a": 1, "an": 1, "and": 1, "aniket": 2, "avishkar": 2, "be": 1, "e": 1, "fname": 1, "gaikwad": 1, "i": 1, "lname": 1, "me": 1, "o": 2, "omkar": 1, "pratik": 1, "shinde": 3, "to": 1, "varun": 1}, ""},
		{"File with no data", "/media/merceadm/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Downloads/file.txt", map[string]int{}, ""},
		{"With non existent file", "/media/merceadm/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Downloads/x.txt", map[string]int{}, "open /media/merceadm/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Downloads/x.txt: no such file or directory"},
	}

	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			actual, err := wordReader(test.filePath)
			assert.DeepEqual(t, test.expected, actual)
			assert.DeepEqual(t, test.err, err)
		})
	}
}
