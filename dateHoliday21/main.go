package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/exp/slices"
	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 21. Write a program to accept a list of holidays (date in any of the supported formats).
* Store this list in an internal array. After the user confirms entering of the holiday list,
* accept a date from the user, and confirm whether it's a working day or not. (All Saturdays
* and Sundays are implicitly considered as holidays). [Arrays & Date Manipulation, 4 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 21
	DATE_FORMAT       = "02/01/2006"
)

func main() {
	var dateToCheckStr string
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	holyList := getHolidays()

	// get date to check from user
	fmt.Println("Enter the date to check (DD/MM/YYYY):")
	fmt.Scanln(&dateToCheckStr)
	dateToCheck, err := time.Parse(DATE_FORMAT, dateToCheckStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1)
	}

	fmt.Println("holiday list ::", holyList)

	if isHoliday(dateToCheck, holyList) {
		fmt.Println(dateToCheckStr, "is a holiday.")
	} else {
		fmt.Println(dateToCheckStr, "is not a holiday.")
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// getHolidays : This function will collect holidays from user
func getHolidays() []time.Time {
	var holidayList []time.Time
	var holidayStr string

	fmt.Println("Enter holiday date (DD/MM/YYYY) or \"proceed\" to end:")
	fmt.Scanln(&holidayStr)
	// loop & add to slice till proceed is entered
	for holidayStr != "proceed" {
		holiday, err := time.Parse(DATE_FORMAT, holidayStr)
		if err != nil {
			log.Fatal(constants.ERROR_17_1)
		}
		holidayList = append(holidayList, holiday)
		fmt.Scanln(&holidayStr)
	}
	return holidayList
}

// isHoliday : This will return true if the given date contains in slice or saturday/sunday
func isHoliday(dateToCheck time.Time, holidayList []time.Time) bool {

	// check for saturday & sunday & holiday in given slice
	if slices.Contains(holidayList, dateToCheck) || dateToCheck.Weekday() == time.Saturday || dateToCheck.Weekday() == time.Sunday {
		return true
	}
	return false
}
