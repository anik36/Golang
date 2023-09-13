package main

import (
	"log"
	"time"

	"redmine.merce.co/git/go-learning-project/anikets/dateFormat14/formats"
)

/******************************************************************************************
 *14. Write a program to print current date/time in following formats (one line per format)
 *   in UTC time e.g. “16 Mar 2022” “Mar 16, 2022” “2022-03-16” “2022-03-16T15:52:00Z”
 *   “Tuesday, 16 March 2022” [Date manipulation, 3 hours]
*******************************************************************************************/

const ASSIGNMENT_NUMBER = 14

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	//for Specific date use
	// date := time.Date(2023, 12, 31, 24, 60, 100, 100, time.Local)

	date := time.Now()
	formats.PrintDateFormats(date)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}
