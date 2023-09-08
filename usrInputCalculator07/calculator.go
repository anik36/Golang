package main

import (
	"errors"
	"fmt"
	"log"
)

/*7. Extend program (5) to accept 3 inputs: “num1”, “num2” and “operation” where operation
could be “+”, “-”, “*” or “/” to represent sum, difference, multiplication or division. The output
will be output of “num1” <operation> “num2”. The output shall be “num1=<num1>
num2=<num2> <operation>=<result>” where “<operation>” will be replaced by the
operation name. Use “sum”, “difference”, “multiply” and “divide” as an operation name when
printing the result. [If/then/else OR switch statement, 3 hours]*/

const ASSIGNMENT_NUMBER = 7
const ERROR = "enter numers only"

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)
	var frstNum, scndNum float64
	var operation int

	fmt.Println("Enter Num1:")
	_, err := fmt.Scan(&frstNum)
	// if not number handle the error
	if err != nil {
		log.Fatal(ERROR)
		return
	}

	fmt.Println("Enter Num2:")
	_, errr := fmt.Scan(&scndNum)
	// if not number handle the error
	if errr != nil {
		log.Fatal(ERROR)
		return
	}

	fmt.Printf("Choose operation to perform then press Enter:\n 1)Sum (+)\n 2)Subtraction (-)\n 3)Multiplication (*)\n 4)Division (/) \n Option :")
	_, eror := fmt.Scan(&operation)
	// if not number handle the error
	if eror != nil {
		log.Fatal(ERROR)
		return
	}

	// do the operation and print the result as per operation
	getResult(frstNum, scndNum, operation)
	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// getResult this function will do the operation as per operation & print the result
func getResult(firstNum float64, secondNum float64, operation int) {
	switch operation {
	case 1:
		fmt.Printf("Num1: %v , Num2: %v, Sum: %.2f \n", firstNum, secondNum, getAddition(firstNum, secondNum))
	case 2:
		fmt.Printf("Num1: %v , Num2: %v, Difference : %.2f \n", firstNum, secondNum, getDiffr(firstNum, secondNum))
	case 3:
		fmt.Printf("Num1: %v , Num2: %v, Multiply : %.2f \n", firstNum, secondNum, getMultiplication(firstNum, secondNum))
	case 4:
		fmt.Printf("Num1: %v , Num2: %v, Divide : %.2f \n", firstNum, secondNum, getDivision(firstNum, secondNum))
	default:
		log.Fatal(errors.New("please choose valid operation"))
	}
}

// getAddition : to add firstNum and secondNum
func getAddition(firstNum float64, secondNum float64) float64 {
	result := firstNum + secondNum
	return result
}

// getDiffr : to subtract secondNum from firstNum
func getDiffr(firstNum float64, secondNum float64) float64 {
	result := firstNum - secondNum
	return result
}

// getMultiplication : to multiply firstNum by secondNum
func getMultiplication(firstNum float64, secondNum float64) float64 {
	result := firstNum * secondNum
	return result
}

// getDivision : to divide firstNum by secondNum
func getDivision(firstNum float64, secondNum float64) float64 {
	result := firstNum / secondNum
	return result
}
