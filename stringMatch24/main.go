package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/exp/slices"
)

/******************************************************************************************
* 24. Write a program to take the names of candidates as input. Prompt user to keep entering
* new names till user enters “done”. Once a list of names are accepted, prompt the user for a
* name. Output shall be “<name> exists” or “<name> does not exist”. A name exists if the
* name exactly matches one of the names provided earlier. Use case insensitive match for
* comparison. [Hash data structure & String operations, 6 hours]
******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 24
	DONE              = "done"
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)
	var name string
	nameRecords := make([]string, 0)

	fmt.Println("nameRecords - len :: ", len(nameRecords), " size/cap :: ", cap(nameRecords))
	fmt.Println("Enter the name to add in records or \"done\" to exit ::")
	fmt.Scanln(&name)
	nameRecords = addRecord(name, nameRecords)

	fmt.Println("Enter the name to check in records ::")
	fmt.Scanln(&name)

	// check if given name contain in slice of records & print the result
	if slices.Contains(nameRecords, name) {
		fmt.Println(name, "exists")
	} else {
		fmt.Println(name, "does not exist")
	}

	fmt.Println("nameRecords - len :: ", len(nameRecords), " size/cap :: ", cap(nameRecords))

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// addRecord : this is to add names into the slice & will return the slice of names given by user
func addRecord(name string, nameRecords []string) []string {
	name = strings.ToLower(strings.TrimSpace(name))
	if name != DONE {
		nameRecords = append(nameRecords, name)
		fmt.Scanln(&name)
		return addRecord(name, nameRecords)
	}
	return nameRecords
}
