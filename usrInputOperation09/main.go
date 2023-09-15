package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

/*9. Extend (7) to accept statistical operation instead of “proceed”. Valid values for “<operation>”
are count (to count number of valid numbers), mean to calculate arithmetic mean, min to
calculate minimum value (minimum of all numbers input), max for maximum value (maximum
of all numbers input). [Switch & functions, 4 hours]*/

const ASSIGNMENT_NUMBER = 9

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	numbers, operation := getData()

	result := getResult(numbers, operation)
	// print result
	fmt.Printf("%s :: %.2f \n", operation, result)
	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// getData : will take all numbers & operation from user
func getData() ([]float64, string) {
	var numbers []float64
	var operation string
	for {
		fmt.Printf("Enter number or type operation (count/mean/minimum/maximum):\n")
		fmt.Scan(&operation)
		operation = strings.ToLower(strings.TrimSpace(operation))
		// PROCEED detect check
		if operation == "count" || operation == "mean" || operation == "min" || operation == "max" {
			fmt.Println("detected ::", operation)
			break
		}
		inputNumber, err := strconv.ParseFloat(operation, 64)
		// error handling for non-numbers
		if err != nil {
			log.Fatal(errors.New("enter valid numers only or operation (count/mean/minimum/maximum)"))
		}
		numbers = append(numbers, inputNumber)
	}
	return numbers, operation
}

// getResult : this func will call the operation func based userInput operation
func getResult(numbers []float64, operation string) float64 {
	var result float64
	switch operation {
	case "count":
		result = float64(len(numbers))
	case "mean":
		result = mean(numbers)
	case "min":
		result = minimum(numbers)
		if len(numbers) == 0 {
			result = 0
		}
	case "max":
		result = maximum(numbers)
		if len(numbers) == 0 {
			result = 0
		}
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

// minimum : it will return minimum numbers
func minimum(list []float64) float64 {
	sort.Float64s(list)
	return list[0]
}

// maximum : it will return maximum numbers
func maximum(list []float64) float64 {
	sort.Float64s(list)
	return list[len(list)-1]
}
