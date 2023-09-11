package main

import (
	"testing"
	"time"

	"gotest.tools/assert"
	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

func TestCalculateHolidays(t *testing.T) {
	holyList := []time.Time{
		time.Date(2023, 9, 14, 0, 0, 0, 0, time.Local), // bro birthday
	}

	testData := []struct {
		description      string
		firstDate        time.Time
		secondDate       time.Time
		expectedHolidays int
		expectedWorkDays int
		err              string
	}{
		{"Valid dates", time.Date(2023, 9, 1, 0, 0, 0, 0, time.Local), time.Date(2023, 9, 18, 0, 0, 0, 0, time.Local), 7, 11, ""},
		{"First date after second date", time.Date(2023, 9, 3, 0, 0, 0, 0, time.Local), time.Date(2023, 9, 2, 0, 0, 0, 0, time.Local), -1, -1, constants.ERROR_22_1},
		{"Holiday in the middle", time.Date(2023, 9, 13, 0, 0, 0, 0, time.Local), time.Date(2023, 9, 15, 0, 0, 0, 0, time.Local), 1, 2, ""},
		{"Both same date", time.Date(2023, 9, 13, 0, 0, 0, 0, time.Local), time.Date(2023, 9, 13, 0, 0, 0, 0, time.Local), -1, -1, constants.ERROR_22_1},
		{"Saturday and Sunday", time.Date(2023, 9, 9, 0, 0, 0, 0, time.Local), time.Date(2023, 9, 10, 0, 0, 0, 0, time.Local), 2, 0, ""},

	}
	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			actualHolidays, actualWorkDays, err := calculateHolidays(test.firstDate, test.secondDate, holyList)
			assert.Equal(t, test.expectedHolidays, actualHolidays)
			assert.Equal(t, test.expectedWorkDays, actualWorkDays)
			assert.Equal(t, test.err, err)
		})
	}
}
