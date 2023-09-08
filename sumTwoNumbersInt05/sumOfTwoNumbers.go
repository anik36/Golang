package main

import (
	"fmt"
	"log"
)

/*5. Write a program to prompt the user for 2 inputs: “num1” and “num2” and generate a sum of
two numbers as output. The program must accept only whole numbers (positive or negative)
or throw an error. The output shall be “num1=<num1> num2=<num2> sum=<result>” where
“<num1>”, “<num2>” and “<result>” will be replaced with actual value. [Basic expressions,
2 hours]*/

const ASSIGNMENT_NUMBER = 5
const ERROR = "enter whole numers only"

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	var frstNum, scndNum int

	fmt.Println("Enter Num1::")
	_, err := fmt.Scan(&frstNum)
	// if not integer handle the error
	if err != nil {
		log.Fatal(ERROR)
		return
	}

	fmt.Println("Enter Num2::")
	_, errr := fmt.Scan(&scndNum)
	// if not integer handle the error
	if errr != nil {
		log.Fatal(ERROR)
		return
	}

	sum := frstNum + scndNum
	// printing final output
	fmt.Printf("Num1: %v , Num2: %v, Sum: %v \n", frstNum, scndNum, sum)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}
