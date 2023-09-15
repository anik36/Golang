package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/exp/slices"
	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 23. Same as (22) to accept a list of holidays, and then prompt a user for two inputs: input date
* as a first argument and a number of business days as a second argument. Number of
* business days will be a positive or negative whole number. The output shall be the date
* relative to an input date +/- the number of business days. Holidays must be excluded while
* calculating the output date. [Date manipulation & loop & boolean expressions, 4 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 23
	DATE_FORMAT       = "02/01/2006"
)

func main() {
	var dateStr string
	var businessDays int
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	holyList := getHolidays()

	fmt.Println("Enter the start date (DD/MM/YYYY):")
	fmt.Scanln(&dateStr)
	startDate, err := time.Parse(DATE_FORMAT, dateStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1)
	}

	fmt.Println("Enter number of business days:")
	fmt.Scanln(&businessDays)
	fmt.Println("businessDays ::", businessDays)

	nextWorkDay := getNextWorkDay(startDate, businessDays, holyList)

	fmt.Println("holiday list ::", holyList)

	fmt.Println("The the next working day is:", nextWorkDay.Format(DATE_FORMAT))

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// getHolidays : will collect holidays from user
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

// getNextWorkDay : it will return the next working day by calculating given business days
func getNextWorkDay(startDate time.Time, givendays int, holyList []time.Time) time.Time {
	var counter, iteration int
	for counter != givendays {
		fmt.Println(iteration)
		iteration++
		// if minus bussines days given
		if givendays <= 0 {
			startDate = startDate.AddDate(0, 0, -1)
			if startDate.Weekday() == time.Sunday {
				startDate = startDate.AddDate(0, 0, -1)
			} else if !(slices.Contains(holyList, startDate) || startDate.Weekday() == time.Saturday) {
				counter--
			}
			// if positive bussiness given
		} else {
			startDate = startDate.AddDate(0, 0, 1)
			if startDate.Weekday() == time.Saturday {
				startDate = startDate.AddDate(0, 0, 1)
			} else if !(slices.Contains(holyList, startDate) || startDate.Weekday() == time.Sunday) {
				counter++
			}
		}
	}
	return startDate
}
