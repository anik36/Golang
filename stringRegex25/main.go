package main

import (
	"fmt"
	"log"
	"regexp"
)

/******************************************************************************************
* 25. Write a program to take the names of candidates as input. Prompt user to keep entering new
* names till user enters “done”. Once a list of names are accepted, prompt the user for a
* search pattern (regex). Output shall be a list of all names where the search pattern exists.
* [Array & Loop & Regex, 8-12 hours]
******************************************************************************************/
const (
	ASSIGNMENT_NUMBER = 25
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	// Create a slice to store the names of candidates.
	var names []string
	var name, searchPattern string

	// Prompt the user to enter names.
	fmt.Println("Enter names of candidates (enter 'done' to stop):")
	for {
		fmt.Scanln(&name)

		if name == "done" {
			break
		}
		names = append(names, name)
	}

	// Prompt the user for a search pattern.
	fmt.Println("Enter a pattern to search :")
	fmt.Scanln(&searchPattern)

	matchedNames := getMatchedNames(searchPattern, names)

	// Print the list of matched names.
	if len(matchedNames) == 0 {
		fmt.Println("No names matched from given pattern.")
	} else {
		fmt.Println("The following names matched from given pattern :")
		fmt.Println(matchedNames)
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// getMatchedNames : returns []string whether the given []string names
// contains any match of the regular expression has in searchPattern (case sensitive).
func getMatchedNames(searchPattern string, names []string) []string {
	regex := regexp.MustCompile(searchPattern)
	var matchedNames []string
	// skip the index of slice and look for value if it match the regex
	// then add it to the matchedNames slice
	for _, name := range names {
		if regex.MatchString(name) {
			matchedNames = append(matchedNames, name)
		}
	}
	return matchedNames
}
