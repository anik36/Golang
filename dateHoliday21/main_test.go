package main

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestIsHoliday(t *testing.T) {
	holidayList := []time.Time{
		time.Date(2023, 12, 25, 0, 0, 0, 0, time.Local), // Christmas
	}

	testData := []struct {
		description    string
		dateToCheck    time.Time
		expectedResult bool
	}{
		{"Holiday in slice", time.Date(2023, 12, 25, 0, 0, 0, 0, time.Local), true},
		{"Saturday", time.Date(2023, 8, 12, 0, 0, 0, 0, time.Local), true},
		{"Sunday", time.Date(2023, 8, 13, 0, 0, 0, 0, time.Local), true},
		{"Not a holiday", time.Date(2023, 8, 8, 0, 0, 0, 0, time.Local), false},
	}

	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			actual := isHoliday(test.dateToCheck, holidayList)
			assert.Equal(t, test.expectedResult, actual)
		})
	}
}
