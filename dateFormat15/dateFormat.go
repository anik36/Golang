package main

import (
	"fmt"
	"log"
	"time"
)

/******************************************************************************************
* 15. Extend (14) to print time in IST timezone. [Date manipulation, 2 hours]
*******************************************************************************************/

const ASSIGNMENT_NUMBER = 15

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	date := time.Now()

	printDateFormats(date)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// printDateFormats : will print date & time in various formats
func printDateFormats(date time.Time) {
	fmt.Println("default ::", date)

	fmt.Println("")
	fmt.Println("1) Format :: 16 Mar 2022")
	fmt.Printf("type1 : %v %v %v \n", date.Day(), date.Month(), date.Year())
	fmt.Printf("type2 : %v \n", date.Format("02 Jan 2006"))

	fmt.Println("")
	fmt.Println("2) Format :: Mar 16, 2022")
	fmt.Printf("type1 : %v %v, %v \n", date.Month(), date.Day(), date.Year())
	fmt.Printf("type2 : %v \n", date.Format("Jan 02, 2006"))

	fmt.Println("")
	fmt.Println("3) Format :: 2022-03-16")
	fmt.Printf("type1 : %v-%v-%v \n", date.Year(), date.Month(), date.Day())
	fmt.Printf("type2 : %v \n", date.Format("2006-Jan-02"))

	fmt.Println("")
	fmt.Println("4) Format :: 2022-03-16T15:52:00Z")
	fmt.Printf("type1 UTC: %v \n", date.UTC())
	fmt.Printf("type2 IST: %v \n", date.Format("2006-01-02T15:04:05 -0700 MST"))

	fmt.Println("")
	fmt.Println("5) Format :: Tuesday, 16 March 2022")
	fmt.Printf("type1 : %v, %v %v %v \n", date.Weekday(), date.Day(), date.Month(), date.Year())
	fmt.Printf("type2 : %v \n", date.Format("Monday, 02 January 2006"))

	fmt.Println("")
	fmt.Println("6) Format :: 11:00PM")
	fmt.Printf("type1 constant: %v \n", date.Format(time.Kitchen))
	fmt.Printf("type2 24-Hrs: %v \n", date.Format("15:04PM"))
	fmt.Printf("type2 12-Hrs: %v \n", date.Format("3:04PM"))

	fmt.Println("")
	fmt.Println("7) Format :: 16/03/2022")
	fmt.Printf("type1 : %v \n", date.Format("02/01/2006"))

	fmt.Println("")
	fmt.Println("8) Format :: IST timezone")
	fmt.Printf("type1 : %v \n", date.Local())
	zone, _ := time.LoadLocation("Asia/Kolkata")
	fmt.Printf("type1 : %v \n", date.In(zone))

}
