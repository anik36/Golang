package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
)

/******************************************************************************************
* 11. Extend (9) to support “countodd” and “counteven” operations to respectively print
* number of times odd numbers and number of even numbers found within the list.
* [Expressions, 2 hours]
*******************************************************************************************/

const PROCEED, ASSIGNMENT_NUMBER = "proceed", 11

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	var numbers []float64
	var operation string

	for {
		fmt.Printf("Enter number or type %v :\n", PROCEED)
		fmt.Scanln(&operation)

		operation = strings.ToLower(operation)
		// end loop if PROCEED detected
		if operation == PROCEED {
			fmt.Println("detected ::", operation)
			break
		}

		num, err := strconv.ParseFloat(operation, 64)
		// error handling for non-numbers
		if err != nil {
			log.Println(constants.ERROR_11_1)
		} else {
			numbers = append(numbers, num)
		}
	}

	fmt.Printf("Enter operation from (count, mean, min, max, countodd, counteven):\n")
	fmt.Scan(&operation)

	operation = strings.ToLower(operation)

	printResult(numbers, operation)

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// printResult : this func will call the operation func based on operation
func printResult(numbers []float64, operation string) {
	var result float64

	sort.Float64s(numbers)

	switch operation {
	case "count":
		result = float64(len(numbers))
	case "mean":
		result = mean(numbers)
	case "min":
		result = numbers[0]
	case "max":
		result = numbers[len(numbers)-1]
	case "countodd":
		oddSlice := getEvenOdd(numbers, "countodd")
		fmt.Printf("Odd Numbers :: %v \n Count :: %v \n", oddSlice, len(oddSlice))
		return
	case "counteven":
		evenSlice := getEvenOdd(numbers, "counteven")
		fmt.Printf("Even Numbers :: %v \nCount :: %v \n", evenSlice, len(evenSlice))
		return
	default:
		log.Fatal(constants.ERROR_11_2, operation)
	}

	fmt.Println(operation, "::", result)
}

// getEvenOdd : this func will return slice based on operation odd/even
func getEvenOdd(numbers []float64, operation string) []float64 {
	var evenNums []float64
	var oddNums []float64

	for _, number := range numbers {
		if math.Remainder(number, 2) == 0 {
			evenNums = append(evenNums, number)
		} else {
			oddNums = append(oddNums, number)
		}
	}
	// falltrough 
	switch operation {
	case "counteven":
		return evenNums
	default:
		return oddNums
	}
}

func mean(list []float64) float64 {
	var result float64
	if len(list) == 0 {
		return 0
	}
	for i := 0; i < len(list); i++ {
		result += list[i]
	}
	return result / float64(len(list))
}
