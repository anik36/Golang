package main

import (
	"fmt"
	"log"
	"time"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 20. Write a program to accept two dates and print the count of week-end days (Consider
* Saturdays and Sundays as week-ends).
*[Loops & Date manipulation & simple expressions, 6 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 20
	DATE_FORMAT       = "02/01/2006"
)

func main() {
	var firstDateStr, secondDateStr string
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	// get the two dates from the user.
	fmt.Println("Enter the first date (DD/MM/YYYY):")
	fmt.Scanln(&firstDateStr)
	firstDate, err := time.Parse(DATE_FORMAT, firstDateStr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter the second date (DD/MM/YYYY):")
	fmt.Scanln(&secondDateStr)
	secondDate, err := time.Parse(DATE_FORMAT, secondDateStr)
	if err != nil {
		log.Fatal(err)
	}

	// calculate the Weekend Days
	weekendDays := calculateWeekendDays(firstDate, secondDate)

	// print the number of weekend days.
	fmt.Println("The number of weekend days is:", weekendDays)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// calculateWeekendDays : it will calculate the weekend days between 2 dates & return count
func calculateWeekendDays(firstDate, secondDate time.Time) int {
	var weekendDays int
	if firstDate.After(secondDate) {
		log.Fatal(constants.ERROR_20_1)
	}
	// add to weekendDays if it is saturday / sunday
	for day := firstDate; day.Before(secondDate) || day.Equal(secondDate); day = day.AddDate(0, 0, 1) {
		if day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
			weekendDays++
		}
	}
	return weekendDays
}
