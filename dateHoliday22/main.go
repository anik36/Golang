package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/exp/slices"
	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 22. Same as (21) to accept a list of holidays, and then prompt a user for two dates (in the
* supported format) and print the number of working days between two dates. Consider both
* dates during the calculation. [Date manipulation & loop & boolean expressions, 4 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 22
	DATE_FORMAT       = "02/01/2006"
)

func main() {
	var firstDateStr, secondDateStr string
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	// get holiday list
	holyList := getHolidays()

	// get the two dates from the user.
	fmt.Println("Enter the first date (DD/MM/YYYY):")
	fmt.Scanln(&firstDateStr)
	firstDate, err := time.Parse(DATE_FORMAT, firstDateStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1)
	}

	fmt.Println("Enter the second date (DD/MM/YYYY):")
	fmt.Scanln(&secondDateStr)
	secondDate, err := time.Parse(DATE_FORMAT, secondDateStr)
	if err != nil {
		log.Fatal(constants.ERROR_17_1)
	}

	holidays, workDays, errMsg := calculateHolidays(firstDate, secondDate, holyList)
	if errMsg != "" {
		log.Fatal(errMsg)
	}
	// Print the number of holidays days.
	fmt.Printf("The number of holidays days from %v to %v are \"%v\". \n", firstDateStr, secondDateStr, holidays)

	// Print the number of workDays days.
	fmt.Printf("The number of workDays days from %v to %v are \"%v\". \n", firstDateStr, secondDateStr, workDays)

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

// calculateHolidays : This will return holidays, workDays by checking from sunday, saturday and holiday list
func calculateHolidays(firstDate, secondDate time.Time, holyList []time.Time) (int, int, string) {
	var holidays, workDays int
	// If firstDate is not before secondDate then give return -1,-1
	if !firstDate.Before(secondDate) {
		return -1, -1, constants.ERROR_22_1
	}
	// Loop through the day if it is Saturday then add 1 to current iterating day
	// (so that when next iteration will add 1 more and will start with adding 2 into current day)
	// if it is Sunday/contains in slice just count holiday
	for day := firstDate; day.Before(secondDate.AddDate(0, 0, 1)); day = day.AddDate(0, 0, 1) {

		switch day.Weekday() == time.Saturday {
		case true:
			holidays += 2
			day = day.AddDate(0, 0, 1)
		case !(slices.Contains(holyList, day) || day.Weekday() == time.Sunday):
			holidays += 1
		default:
			workDays++
		}
	}
	return holidays, workDays, ""
}
