package formats

import (
	"fmt"
	"time"
)

// PrintDateFormats : will print date & time in various formats
func PrintDateFormats(date time.Time) {
	fmt.Println("=>  default :", date)

	fmt.Println("")
	fmt.Println("1) Format :: 16 Mar 2022")
	fmt.Printf("type1 dd-mmm-yyyy: %v %v %v \n", date.Day(), date.Month(), date.Year())
	fmt.Printf("type2 dd-mmm-yyyy: %v \n", date.Format("02 Jan 2006"))

	fmt.Println("")
	fmt.Println("2) Format :: Mar 16, 2022")
	fmt.Printf("type1 mmm dd, yyyy: %v %v, %v \n", date.Month(), date.Day(), date.Year())
	fmt.Printf("type2 mmm dd, yyyy: %v \n", date.Format("Jan 02, 2006"))

	fmt.Println("")
	fmt.Println("3) Format :: 2022-03-16")
	fmt.Printf("type1 yyyy-mmm-dd: %v-%v-%v \n", date.Year(), date.Month(), date.Day())
	fmt.Printf("type2 yyyy-mmm-dd: %v \n", date.Format("2006-Jan-02"))

	fmt.Println("")
	fmt.Println("4) Format :: 2022-03-16T15:52:00Z")
	fmt.Printf("type1 UTC: %v \n", date.UTC())
	fmt.Printf("type2 UTC: %v \n", date.Format("2006-01-02T15:04:05Z"))

	fmt.Println("")
	fmt.Println("5) Format :: Tuesday, 16 March 2022")
	fmt.Printf("type1 dd-mmm-yyyy: %v, %v %v %v \n", date.Weekday(), date.Day(), date.Month(), date.Year())
	fmt.Printf("type2 dd-mmm-yyyy: %v \n", date.Format("Monday, 02 January 2006"))

	fmt.Println("")
	fmt.Println("6) Format :: 3:04PM")
	fmt.Printf("type1 constant: %v \n", date.Format(time.Kitchen))
	fmt.Printf("type2 24-Hrs: %v \n", date.Format("15:04PM"))
	fmt.Printf("type2 12-Hrs: %v \n", date.Format("3:04PM"))

	fmt.Println("")
	fmt.Println("7) Format :: 16/03/2022")
	fmt.Printf("type1 dd/mm/yyyy %v \n", date.Format("02/01/2006"))
	fmt.Println("")
}

// PrintMyFormat : this function will print the given date as per flag input
func PrintMyFormat(date time.Time, flag int) {
	switch flag {
	case 1:
		fmt.Println("1) Format :: 16 Mar 2022")
		fmt.Printf("type1 : %v %v %v \n", date.Day(), date.Month(), date.Year())
		fmt.Printf("type2 : %v \n", date.Format("02 Jan 2006"))
	case 2:
		fmt.Println("2) Format :: Mar 16, 2022")
		fmt.Printf("type1 : %v %v, %v \n", date.Month(), date.Day(), date.Year())
		fmt.Printf("type2 : %v \n", date.Format("Jan 02, 2006"))
	case 3:
		fmt.Println("3) Format :: 2022-03-16")
		fmt.Printf("type1 : %v \n", date.Format(time.DateOnly))
		fmt.Printf("type2 : %v \n", date.Format("2006-01-02"))
	case 4:
		fmt.Println("4) Format :: 2022-03-16T15:52:00Z")
		fmt.Printf("type1 UTC: %v \n", date.UTC())
		fmt.Printf("type2 UTC: %v \n", date.Format("2006-01-02T15:04:05 -0700 MST"))
	case 5:
		fmt.Println("5) Format :: Tuesday, 16 March 2022")
		fmt.Printf("type1 : %v, %v %v %v \n", date.Weekday(), date.Day(), date.Month(), date.Year())
		fmt.Printf("type2 : %v \n", date.Format("Monday, 02 January 2006"))
	case 6:
		fmt.Println("6) Format :: 15:04PM")
		fmt.Printf("type1 constant: %v \n", date.Format(time.Kitchen))
		fmt.Printf("type2 24-Hrs: %v \n", date.Format("15:04PM"))
		fmt.Printf("type2 12-Hrs: %v \n", date.Format("3:04PM"))

	default:
		fmt.Println("7) Format :: dd/mm/yyyy")
		fmt.Printf("type1 : %v \n", date.Format("02/01/2006"))
	}
	fmt.Println("")
}
