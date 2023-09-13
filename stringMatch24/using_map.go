package main

import (
	"fmt"
	"log"
	"strings"
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
	mapNameRecords := make(map[string]bool)

	fmt.Println("Enter the name to add in records or \"done\" to exit ::")
	fmt.Scanln(&name)
	mapNameRecords = addToMap(name, mapNameRecords)
	// fmt.Println("RECORD ::", mapNameRecords)

	fmt.Println("Enter the name to check in records ::")
	fmt.Scanln(&name)

	if mapNameRecords[name] {
		fmt.Println(name, "exists in map")
	} else {
		fmt.Println(name, "does not exist in map")
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

func addToMap(name string, mapNameRecords map[string]bool) map[string]bool {
	name = strings.ToLower(strings.TrimSpace(name))
	if name != DONE {
		mapNameRecords[name] = true
		fmt.Scanln(&name)
		return addToMap(name, mapNameRecords)
	}
	return mapNameRecords
}
