package main

import (
	"fmt"
	"log"
	"os"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
	"redmine.merce.co/git/go-learning-project/anikets/fileHtmlDownload27/download"
)

/******************************************************************************************
* 28. Extend (26) to accept URL as a command line argument instead of a hardcoded URL within
* the program. [Revision of previous concepts, 2 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 28
)

// go run main.go "https://remiges.tech"
func main() {
	URL := os.Args[1]
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	result := download.GetHtml(URL, constants.PATH_FILE+constants.HTML_NAME)
	if result {
		fmt.Println("file is downloaded successfully.")
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}
