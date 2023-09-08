package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*10. Extend (8) to support “sort” operation. Use an in-built function call for sorting numbers.*/

const PROCEED, ASSIGNMENT_NUMBER = "proceed", 10

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	numbers, operation := getData()

	result := getResult(numbers, operation)
	if result != 0.0 {
		fmt.Printf("%s :: %.2f \n", operation, result)
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// getData : will take all numbers & operation from user
func getData() ([]float64, string) {
	var numbers []float64
	var operation string
	for {
		fmt.Printf("Enter number or type %v :\n", PROCEED)
		fmt.Scan(&operation)
		operation = strings.ToLower(strings.TrimSpace(operation))
		// PROCEED detect check
		if operation == PROCEED {
			fmt.Println("detected ::", operation)
			break
		}
		inputNumber, err := strconv.ParseFloat(operation, 64)
		// error handling for non-numbers
		if err != nil {
			log.Fatal(errors.New("enter valid numers only or operation"))
		}
		numbers = append(numbers, inputNumber)
	}
	fmt.Printf("Enter operation from (count, mean, min, max, sort):\n")
	_, err := fmt.Scan(&operation)
	if err != nil {
		log.Fatal("enter correct operation!")
	}
	operation = strings.ToLower(strings.TrimSpace(operation))
	return numbers, operation
}

// getResult : this func will return result by performing operation
func getResult(numbers []float64, operation string) float64 {
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
	case "sort":
		fmt.Println("Sorted numbers::", numbers)
	default:
		log.Fatal("operation not supported ::", operation)
	}
	return result
}

// mean : it will return mean of given numbers
func mean(list []float64) float64 {
	var result float64
	for i := 0; i < len(list); i++ {
		result += list[i]
	}
	return result / float64(len(list))
}
