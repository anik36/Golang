package main

import (
	"fmt"
	"log"
	"time"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 17. Write a program to accept two dates (any of the formats supported in the earlier problem)
* and print a difference in human readable format e.g. “1 year 2 day 32 minutes”. [Date
* manipulation, 3 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 17
	DATE_FORMAT       = "02/01/2006T15:04:05"
)

func main() {

	var firstDateStr, secondDateStr string

	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	fmt.Println("Enter first date in format (DD/MM/YYYThh:mm:ss):")
	fmt.Scan(&firstDateStr)

	firstDate, err := time.Parse(DATE_FORMAT, firstDateStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1, err)
	}

	fmt.Println("Enter second date in format (DD/MM/YYYThh:mm:ss):")
	fmt.Scan(&secondDateStr)

	secondDate, err := time.Parse(DATE_FORMAT, secondDateStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1, err)
	}

	fmt.Println("First Date ::", firstDate)
	fmt.Println("Second Date ::", secondDate)

	// calculate years, month, days and time between dates
	year, month, day, hour, min, sec := calculateDifference(firstDate, secondDate)

	fmt.Printf("difference %d years, %d months, %d days, %d hours, %d mins and %d seconds.\n", year, month, day, hour, min, sec)
	fmt.Printf("")

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// calculateDifference : to calculate difference between 2 dates
func calculateDifference(firstDate, secondDate time.Time) (year, month, day, hour, min, sec int) {
	// if firstDate is after secondDate, swap the two dates
	if firstDate.After(secondDate) {
		firstDate, secondDate = secondDate, firstDate
	}

	// get the year, month, and day of the two dates
	year1, month1, day1 := firstDate.Date()
	year2, month2, day2 := secondDate.Date()

	// get the hour, minute, and second of the two dates
	hours1, minute1, second1 := firstDate.Clock()
	hours2, minute2, second2 := secondDate.Clock()

	// calculate the difference
	year = year2 - year1
	month = int(month2 - month1)
	day = day2 - day1
	hour = hours2 - hours1
	min = minute2 - minute1
	sec = second2 - second1

	// fmt.Println("year::", year, " month::", month, " day::", day, " hour::", hour, " mint::", min)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}

	if day < 0 {
		day += 31
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}
