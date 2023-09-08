package main

import (
	"fmt"
	"log"
	"strings"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 13. Write a program to prompt for three inputs: character to be used for display, num1 to
* represent number of rows and num2 to represent number of columns. The output will be a
* rectangular matrix where each cell will print a character input as a first input value.
* [Loops, 4 hours]
******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 13
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	printMatrix(getData())

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// getData : will get the required data from user
func getData() (int, int, string) {
	var char string
	var row, col int

	fmt.Print("Enter a character to be used for display: ")
	// get char to print
	_, errr := fmt.Scanln(&char)
	char = strings.TrimSpace(char)
	// fmt.Printf("%c", char[0])
	if errr != nil || len(char) > 1 {
		log.Fatal(constants.ERROR_13_1)
	}

	fmt.Print("Enter the number of rows: ")
	// get number of rows
	_, err := fmt.Scan(&row)
	if err != nil || row < 0 {
		log.Fatal(constants.ERROR_12_1)
	}

	fmt.Print("Enter the number of columns: ")
	// get number of columns
	_, err1 := fmt.Scan(&col)
	if err1 != nil || col < 0 {
		log.Fatal(constants.ERROR_12_1)
	}
	return row, col, char
}

// printMatrix : this function will print the matrix with given character
func printMatrix(row, col int, character string) {
	// Print a rectangular matrix where each cell will print a character input as a first input value
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print(" ", character)
		}
		fmt.Println()
	}
}
