package main

import (
	"testing"
	"time"
)

/************************************
* Not a Leap year
*************************************/
func TestOne(t *testing.T) {
	date := time.Date(1994, 12, 24, 6, 0, 0, 0, time.UTC)

	expected := false
	got := isLeapYear(date)

	if got != expected {
		t.Errorf("expected : %v , got : %v", expected, got)
	}
}

/************************************
* Leap year
*************************************/
func TestTwo(t *testing.T) {
	date := time.Date(2024, 12, 24, 6, 0, 0, 0, time.UTC)

	expected := true
	got := isLeapYear(date)

	if got != expected {
		t.Errorf("expected : %v , got : %v", expected, got)
	}
}
