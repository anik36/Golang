package main

import (
	"fmt"
	"log"
	"os"
)
const ASSIGNMENT_NUMBER = 3

// go run helloUser.go Aniket
// this program takes name as input from command line
func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	if len(os.Args) > 2 {
		log.Fatal("only enter name while running the program")
	}
	// to accept argument from command line
	name := os.Args[1]
	fmt.Println("Hello", name)
	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}
