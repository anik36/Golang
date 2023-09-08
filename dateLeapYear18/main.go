package main

import (
	"fmt"
	"log"
	"time"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 18. Write a program to accept a date and print whether the date falls in a leap year. Accept a
* date in any format supported in one of the previous problems.
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 18
	DATE_FORMAT       = "02/01/2006"
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	var inputDateStr string
	fmt.Println("Enter date in format (DD/MM/YYY):")
	fmt.Scan(&inputDateStr)

	givenDate, err := time.Parse(DATE_FORMAT, inputDateStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1)
	}

	result := isLeapYear(givenDate.Local())

	if result {
		fmt.Println(inputDateStr, "is a leap year.")
	} else {
		fmt.Println(inputDateStr, "is not a leap year.")
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// isLeapYear : will check if given date year has 366 days and return true as leap year
func isLeapYear(givenDate time.Time) bool {
	var isLeap bool
	// Seperating the year to check
	year, _, _ := givenDate.Date()
	startDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(year, 12, 31, 24, 0, 0, 0, time.Local)

	// Calculate the time difference
	diff := endDate.Sub(startDate)

	days := diff.Hours() / 24

	fmt.Println("Days ::", days)

	if days == 366 {
		isLeap = true
	}
	return isLeap
}
