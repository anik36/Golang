package main

import (
	"reflect"
	"testing"
)

// Test case 1: Input is a slice of all even numbers, operation is "counteven".
func TestOne(t *testing.T) {
	numbers := []float64{0, 2, 4, 6, 8}
	expectedResult := []float64{0, 2, 4, 6, 8}
	actualResult := getEvenOdd(numbers, "counteven")
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Expected result to be %v, got %v", expectedResult, actualResult)
	}
}

// Test case 2: Input is a slice of all odd numbers, operation is "countodd".
func TestTwo(t *testing.T) {
	numbers := []float64{1, 3, 5, 7, 9}
	expectedResult := []float64{1, 3, 5, 7, 9}
	actualResult := getEvenOdd(numbers, "countodd")
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Expected result to be %v, got %v", expectedResult, actualResult)
	}
}

// Test case 3: Input is a mixed slice of even and odd numbers, operation is "counteven".
func TestThree(t *testing.T) {
	numbers := []float64{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	expectedResult := []float64{0, 2, 4, 6, 8}
	actualResult := getEvenOdd(numbers, "counteven")
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Expected result to be %v, got %v", expectedResult, actualResult)
	}
}

// Test case 4: Input is a mixed slice of even and odd numbers, operation is "countodd".
func TestFour(t *testing.T) {
	numbers := []float64{0, 2, 4, 6, 8, 1, 3, 5, 7, 9}
	expectedResult := []float64{1, 3, 5, 7, 9}
	actualResult := getEvenOdd(numbers, "countodd")
	if !reflect.DeepEqual(actualResult, expectedResult) {
		t.Errorf("Expected result to be %v, got %v", expectedResult, actualResult)
	}
}

func TestMeanEmptyList(t *testing.T) {
	list := []float64{}
	mean := mean(list)
	if mean != 0 {
		t.Errorf("Expected mean of empty list to be 0, got %f", mean)
	}
}

func TestMeanOneElementList(t *testing.T) {
	list := []float64{1.0}
	mean := mean(list)
	if mean != 1 {
		t.Errorf("Expected mean of list with one element to be 1, got %f", mean)
	}
}

func TestMeanMultipleElementsList(t *testing.T) {
	list := []float64{1.0, 2.0, 3.0}
	mean := mean(list)
	if mean != 2 {
		t.Errorf("Expected mean of list with multiple elements to be 2, got %f", mean)
	}
}
