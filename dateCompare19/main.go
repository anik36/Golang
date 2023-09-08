package main

import (
	"fmt"
	"log"
	"time"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 19. Write a program to accept two dates (any of the supported period) and print an output
* whether date1 and date2 are equal, date1 is earlier than date2 or date1 is later than date2.
* [Date comparison, 1 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 19
	DATE_FORMAT       = "02/01/2006"
)

func main() {
	var firstDateStr, secondDateStr string
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	// Get the two dates from the user.
	fmt.Println("Enter the first date (DD/MM/YYYY):")
	fmt.Scanln(&firstDateStr)
	firstDate, err := time.Parse(DATE_FORMAT, firstDateStr)
	if err != nil {
		log.Fatalf("%s :: %v", constants.ERROR_17_1, err)
	}

	fmt.Println("Enter the second date (DD/MM/YYYY):")
	fmt.Scanln(&secondDateStr)
	secondDate, err := time.Parse(DATE_FORMAT, secondDateStr)
	if err != nil {
		log.Fatalf("%s :: %v", constants.ERROR_17_1, err)
	}

	// calling date check func
	fmt.Println("Approach 1 ::")
	printIsEqual(firstDate, secondDate)

	fmt.Println("Approach 2 ::")
	printIsEqual2(firstDate, secondDate)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// printIsEqual : will print the result of two date are equal or not
func printIsEqual(firstDate, secondDate time.Time) {
	// Compare the two dates.
	if firstDate.Equal(secondDate) {
		fmt.Println("The two dates are equal.")
	} else if firstDate.Before(secondDate) {
		fmt.Println("first_date is earlier than second_date.")
	} else {
		fmt.Println("first_date is later than second_date.")
	}
}

// printIsEqual2 : will print the result of two date are equal or not
func printIsEqual2(firstDate, secondDate time.Time) {

	diff := firstDate.Sub(secondDate)

	if diff == 0 {
		fmt.Println("The two dates are equal.")
	} else if diff < 0 {
		fmt.Println("first_date is earlier than second_date.")
	} else {
		fmt.Println("first_date is later than second_date.")
	}
}
