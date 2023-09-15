package main

import (
	"fmt"
	"log"
)

/* 6. Write a program similar to (5) above, but accept floating point numbers. [Data types, 2
hours] */

const ASSIGNMENT_NUMBER = 6
const ERROR = "enter numers only"

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	var frstNum, scndNum float64

	fmt.Println("Enter Num1::")
	_, err := fmt.Scan(&frstNum)
	// if not number handle the error
	if err != nil {
		log.Fatal(ERROR)
		return
	}

	fmt.Println("Enter Num2::")
	_, errr := fmt.Scan(&scndNum)
	// if not number handle the error
	if errr != nil {
		log.Fatal(ERROR)
		return
	}

	sum := frstNum + scndNum

	// printing final output
	fmt.Printf("Num1: %v , Num2: %v, Sum: %v \n", frstNum, scndNum, sum)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}
