package main

import (
	"testing"
	"time"

	"gotest.tools/assert"
)

func TestGetNextWorkDay(t *testing.T) {
	holyList := []time.Time{
		time.Date(2023, 8, 15, 0, 0, 0, 0, time.Local), // Independence Day
		time.Date(2023, 8, 16, 0, 0, 0, 0, time.Local), // My Holiday
	}

	testData := []struct {
		description  string
		startDate    time.Time
		businessDays int
		expectedDate time.Time
	}{
		{"Positive business days", time.Date(2023, 8, 6, 0, 0, 0, 0, time.Local), 2, time.Date(2023, 8, 8, 0, 0, 0, 0, time.Local)},
		{"Negative business  days", time.Date(2023, 8, 6, 0, 0, 0, 0, time.Local), -2, time.Date(2023, 8, 3, 0, 0, 0, 0, time.Local)},
		{"Zero business days", time.Date(2023, 8, 28, 0, 0, 0, 0, time.Local), 0, time.Date(2023, 8, 28, 0, 0, 0, 0, time.Local)},
		{"Holiday in between", time.Date(2023, 8, 11, 0, 0, 0, 0, time.Local), 3, time.Date(2023, 8, 18, 0, 0, 0, 0, time.Local)},
	}

	for _, test := range testData {
		t.Run(test.description, func(t *testing.T) {
			actual := getNextWorkDay(test.startDate, test.businessDays, holyList)
			assert.Equal(t, test.expectedDate, actual)
		})
	}
}

/*************************************************************************************
* 1st Approach : either you can loop all data like above else
* you can write seperate methode like below for all
* ************************************************************************************/

// t.Run("Negative business  days", func(t *testing.T) {
// 	firstDate := time.Date(2023, 8, 6, 0, 0, 0, 0, time.Local)
// 	givenDays := -2
// 	expected := time.Date(2023, 8, 3, 0, 0, 0, 0, time.Local)
// 	actual := getNextWorkDay(firstDate, givenDays, holyList)
// 	assert.Equal(t, expected, actual)
// })

// t.Run("Zero business days", func(t *testing.T) {
// 	firstDate := time.Date(2023, 8, 28, 0, 0, 0, 0, time.Local)
// 	givenDays := 0
// 	expected := time.Date(2023, 8, 28, 0, 0, 0, 0, time.Local)
// 	actual := getNextWorkDay(firstDate, givenDays, holyList)
// 	assert.Equal(t, expected, actual)
// })

// t.Run("Holiday in between", func(t *testing.T) {
// 	firstDate := time.Date(2023, 8, 11, 0, 0, 0, 0, time.Local)
// 	givenDays := 3
// 	expected := time.Date(2023, 8, 18, 0, 0, 0, 0, time.Local)
// 	actual := getNextWorkDay(firstDate, givenDays, holyList)
// 	assert.Equal(t, expected, actual)
// })
