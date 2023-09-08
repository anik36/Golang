package main

import (
	"testing"
	"time"
)

/******************************************
*  1 = America/Toronto
******************************************/
func TestOne(t *testing.T) {
	date := time.Date(1994, 12, 24, 6, 0, 0, 0, time.UTC)

	timeZoneChoice := 1
	expectedLoc := "America/Toronto"
	zone, _ := time.LoadLocation(expectedLoc)
	expectedDate := date.In(zone)

	gotDate, gotLocation := getDateByLocation(date, timeZoneChoice)

	if expectedLoc != gotLocation {
		t.Errorf("expected location : %v , got location : %v", expectedLoc, gotLocation)
	}
	if expectedDate.GoString() != gotDate.GoString() {
		t.Errorf("expected date : %v , got date : %v", expectedDate, gotDate)
	}
}

/******************************************
*  2 = Asia/Qatar
******************************************/
func TestTwo(t *testing.T) {
	date := time.Date(1994, 12, 24, 6, 0, 0, 0, time.UTC)

	timeZoneChoice := 2
	expectedLoc := "Asia/Qatar"
	zone, _ := time.LoadLocation(expectedLoc)
	expectedDate := date.In(zone)

	gotDate, gotLocation := getDateByLocation(date, timeZoneChoice)

	if expectedLoc != gotLocation {
		t.Errorf("expected location : %v , got location : %v", expectedLoc, gotLocation)
	}
	if expectedDate.GoString() != gotDate.GoString() {
		t.Errorf("expected date : %v , got date : %v", expectedDate, gotDate)
	}
}

/******************************************
*  7 = Europe/Istanbul
******************************************/
func TestThree(t *testing.T) {
	date := time.Date(1994, 12, 24, 6, 0, 0, 0, time.UTC)

	timeZoneChoice := 7
	expectedLoc := "Europe/Istanbul"
	zone, _ := time.LoadLocation(expectedLoc)
	expectedDate := date.In(zone)

	gotDate, gotLocation := getDateByLocation(date, timeZoneChoice)

	if expectedLoc != gotLocation {
		t.Errorf("expected location : %v , got location : %v", expectedLoc, gotLocation)
	}
	if expectedDate.GoString() != gotDate.GoString() {
		t.Errorf("expected date : %v , got date : %v", expectedDate, gotDate)
	}
}

/******************************************
*  0 = nil
* default = Australia/Sydney
******************************************/
func TestFour(t *testing.T) {
	date := time.Date(1994, 12, 24, 6, 0, 0, 0, time.UTC)

	timeZoneChoice := 0
	expectedLoc := "Australia/Sydney"
	zone, _ := time.LoadLocation(expectedLoc)
	expectedDate := date.In(zone)

	gotDate, gotLocation := getDateByLocation(date, timeZoneChoice)

	if expectedLoc != gotLocation {
		t.Errorf("expected location : %v , got location : %v", expectedLoc, gotLocation)
	}
	if expectedDate.GoString() != gotDate.GoString() {
		t.Errorf("expected date : %v , got date : %v", expectedDate, gotDate)
	}
}
