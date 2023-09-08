package main

import (
	"testing"
	"time"
)

func TestDateWithTime(t *testing.T) {

	firstDate, _ := time.Parse(DATE_FORMAT, "24/12/1994T06:30:00")
	secondDate, _ := time.Parse(DATE_FORMAT, "10/08/2023T03:00:15")

	yearTest, monthTest, dayTest, hourTest, minTest, secTest := calculateDifference(firstDate, secondDate)

	year, month, day, hour, min, sec := 28, 7, 16, 20, 30, 15

	if yearTest != year {
		t.Errorf("Year faild :: got %d, wanted %d", yearTest, year)
	}

	if monthTest != month {
		t.Errorf("Month faild :: got %d, wanted %d", monthTest, month)
	}

	if dayTest != day {
		t.Errorf("Day faild :: got %d, wanted %d", dayTest, day)
	}

	if hourTest != hour {
		t.Errorf("Hour faild :: got %d, wanted %d", hourTest, hour)
	}

	if minTest != min {
		t.Errorf("Min faild :: got %d, wanted %d", minTest, min)
	}

	if secTest != sec {
		t.Errorf("Min faild :: got %d, wanted %d", secTest, sec)
	}
}

func TestDateOnly(t *testing.T) {

	firstDate, _ := time.Parse(DATE_FORMAT, "24/12/1994T00:00:00")
	secondDate, _ := time.Parse(DATE_FORMAT, "10/08/2023T00:00:00")

	yearTest, monthTest, dayTest, hourTest, minTest, secTest := calculateDifference(firstDate, secondDate)

	year, month, day, hour, min, sec := 28, 7, 17, 0, 0, 0

	if yearTest != year {
		t.Errorf("Year faild :: got %d, wanted %d", yearTest, year)
	}

	if monthTest != month {
		t.Errorf("Month faild :: got %d, wanted %d", monthTest, month)
	}

	if dayTest != day {
		t.Errorf("Day faild :: got %d, wanted %d", dayTest, day)
	}

	if hourTest != hour {
		t.Errorf("Hour faild :: got %d, wanted %d", hourTest, hour)
	}

	if minTest != min {
		t.Errorf("Min faild :: got %d, wanted %d", minTest, min)
	}

	if secTest != sec {
		t.Errorf("Min faild :: got %d, wanted %d", secTest, sec)
	}
}
