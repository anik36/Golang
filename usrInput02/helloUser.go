package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const ASSIGNMENT_NUMBER = 2

func main() {

	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	// 1st approach using "bufio" package

	// creating reader
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("1. Enter Your Name:")
	// keep reading till hit the new line
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello", input)

	// 2nd approach using "fmt" package
	var name string
	fmt.Println("2. Enter Your Name:")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}
