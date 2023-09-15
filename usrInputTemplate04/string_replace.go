package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const ASSIGNMENT_NUMBER = 4

const NAMETAG = "<name>"

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	if len(os.Args) > 3 {
		log.Fatal("only enter template & name while running the program")
	}

	template := os.Args[1]
	name := os.Args[2]
	// error if no NAMETAG
	if !nameTagCheck(template) {
		log.Fatal("template must contain", NAMETAG)
	}
	// error if blank name input
	if strings.TrimSpace(name) == "" {
		log.Fatal("please enter your name")
	}

	result := strings.ReplaceAll(template, NAMETAG, name)
	fmt.Println(strings.ReplaceAll(result, "\n", ""))

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}

// nameTagCheck : is to check given string contain a NAMETAG
func nameTagCheck(template string) bool {
	result := strings.Contains(template, NAMETAG)
	return result
}
