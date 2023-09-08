package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"redmine.merce.co/git/go-learning-project/anikets/regex"
)

/******************************************************************************************
* 29. Write a program to accept a filename from command line argument, read a file and print the
* number of times each word occurs in a file. Perform case insensitive match while counting
* the occurrence of each word. [Hashmap & File operation & String operation, 6 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 29
)

// go run wordCount.go "/media/merceadm/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Downloads/input.txt"
func main() {
	args := os.Args[1:]
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	if len(args) != 1 {
		log.Fatal("Please provide the filename as the first argument")
	}

	filename := args[0]

	wordCounts, err := wordReader(filename)
	if err != "" {
		log.Fatal(err)
	}

	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}

func wordReader(filePath string) (map[string]int, string) {
	wordCounts := make(map[string]int)
	file, err := os.Open(filePath)
	if err != nil {
		return wordCounts, fmt.Sprint(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		// remove sysmbols from word
		word = regex.RemoveSymbols(word)
		wordCounts[word]++
	}
	return wordCounts, ""
}
