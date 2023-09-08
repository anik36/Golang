package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/******************************************************************************************
* 30. Extend the above program to ignore common words (e.g. “the”, “a”, “an”, …) and single letter
* words (e.g. “I”). [Revision of previous concepts, 3 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 30
)

// go run main.go "/media/merceadm/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Downloads/input.txt"

func main() {
	commonWords := []string{"the", "an", "to", "of", "and", "in", "on", "is", "it"}

	args := os.Args[1:]
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	if len(args) != 1 {
		log.Fatal("Please provide the filename as the first argument")
	}

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCounts := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) < 2 || isCommonWord(word, &commonWords) {
			continue
		}
		wordToAdd := strings.ToLower(word)

		wordCounts[wordToAdd]++
	}

	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

// isCommonWord checks the given string is belongs to given slice or not & return bool value.
func isCommonWord(word string, commonWords *[]string) bool {
	for _, commonWord := range *commonWords {
		if word == commonWord {
			return true
		}
	}
	return false
}
