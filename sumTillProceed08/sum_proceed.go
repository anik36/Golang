package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
8. Write a program to prompt for a sequence of numbers, one number at a time, till the user
enters “proceed”. Upon receiving “proceed”, the program shall calculate the sum of all
numbers and produce an output. Ensure that only valid numbers are considered as an input;
if you detect an invalid number input, give a meaningful error message and exit. Your
program must work correctly even if the user enters zero or just one number. [Data types
and validations & for-loop, 4 hours]
*/

const PROCEED, ASSIGNMENT_NUMBER = "proceed", 8

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)
	// calculate sum
	sum := calculateSum()
	// printing the sum
	fmt.Println("Sum:", sum)
	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// calculateSum will do the sum by getting input from user
func calculateSum() float64 {
	var sum float64
	// creating reader
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Enter a number or %s to stop:\n", PROCEED)
		// keep reading till hit the new line
		inputString, _ := reader.ReadString('\n')
		// break the loop if detected
		if strings.TrimSpace(inputString) == PROCEED {
			log.Println("detected ::", PROCEED)
			break
		}

		// parse string to float for Num1
		numberToAdd, err := strconv.ParseFloat(strings.TrimSpace(inputString), 64)
		// error if not a number or parse failed
		if err != nil {
			log.Fatal(errors.New("enter proper one number at a time or proceed"))
		}
		// add that num
		sum += numberToAdd
	}
	return sum
}
