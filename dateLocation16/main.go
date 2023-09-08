package main

import (
	"fmt"
	"log"
	"time"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
	"redmine.merce.co/git/go-learning-project/anikets/dateFormat14/formats"
)

/******************************************************************************************
* 16. Extend (14) to print supported timezones, and accept a valid timezone as input and print
* time as per the time zone selected by an end user. [Date manipulation & switch
* statement, 3 hours]
*******************************************************************************************/

const ASSIGNMENT_NUMBER = 16

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)
	var timeZone int

	// seting the custom date
	// date := time.Date(1994, 12, 24, 6, 0, 0, 0, time.UTC)
	date := time.Now()

	fmt.Println("========== Which Timezone would you like to check? ==========\n1) America \n2) Qatar \n3) Shanghai \n4) Macau \n5) Singapore \n6) India \n7) Turkey \n8) Australia ")
	fmt.Print(">")

	_, err := fmt.Scanln(&timeZone)
	if err != nil || timeZone > 8 || 1 > timeZone {
		log.Fatal(constants.ERROR_16_1)
	}

	// This function gets the new date in the user's selected timezone.
	newDate, location := getDateByLocation(date, timeZone)

	fmt.Printf("Time in %v :: %v \n \n", location, newDate)

	fmt.Println("========== Priting Format From Assingment 14 using new Date ==========")
	fmt.Println("Which format would you like to check? \n1) 16 Mar 2022 \n2) Mar 16, 2022 \n3) 2022-03-16 \n4) 2022-03-16T15:52:00Z \n5) Tuesday, 16 March 2022 \n6) 11:00PM \n7) 16/03/2022 ")
	fmt.Print(">")

	_, errr := fmt.Scanln(&timeZone)
	if errr != nil || timeZone > 7 || 1 > timeZone {
		log.Fatal(constants.ERROR_16_1)
	}

	formats.PrintMyFormat(newDate, timeZone)
	fmt.Println("=====================================================================")

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// getDateByLocation : This function returns the new date in the user's selected timezone.
func getDateByLocation(date time.Time, userChoice int) (time.Time, string) {
	var zone *time.Location
	var locationStr string

	// This switch statement gets the timezone and location for the user's selected choice.
	switch userChoice {
	case 1:
		locationStr = "America/Toronto"
	case 2:
		locationStr = "Asia/Qatar"
	case 3:
		locationStr = "Asia/Shanghai"
	case 4:
		locationStr = "Asia/Macau"
	case 5:
		locationStr = "Asia/Singapore"
	case 6:
		locationStr = "Asia/Kolkata"
	case 7:
		locationStr = "Europe/Istanbul"
	default:
		locationStr = "Australia/Sydney"
	}
	zone, _ = time.LoadLocation(locationStr)
	return date.In(zone), locationStr
}
