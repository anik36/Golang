package main

import (
	"fmt"
	"log"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 12. Write a program to prompt for two whole positive numbers -- “num1” and “num2”. Print
* multiplication table for the number e.g. for num1=3 and num2=20, output will be “3 * 1 = 3\n3
* 2 = 6\n … \n3 * 20 = 60”. [For loop, 6 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 12
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	numOne, numTwo := getNumber(), getNumber()

	printTable(numOne, numTwo)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// getNumber : will get the number from user else gives error
func getNumber() int {
	var number int
	fmt.Println("Enter number:")
	_, err := fmt.Scanln(&number)
	// error if not an integer
	if err != nil || number < 0 {
		log.Fatal(constants.ERROR_12_1)
	}
	return number
}

// printTable : will print the multiplication table
func printTable(numOne, numTwo int) {
	for i := 1; i <= numTwo; i++ {
		// %.2 to take 2 digit into consideration
		// %2 to take space of 2 digit into consideration
		fmt.Printf("%.2d * %2d = %.2d \n", numOne, i, (numOne * i))
	}
}
