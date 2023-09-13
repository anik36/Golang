package main

import (
	"fmt"
	"log"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
	"redmine.merce.co/git/go-learning-project/anikets/fileHtmlDownload27/download"
)

/******************************************************************************************
* 27. Extend (26) to use URL classes instead of a browser to download a file. Rest of the
* functionalities will be the same as the previous problem. [Java classes & file operations,
* 8-12 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 27
	URL               = "https://merce.co"
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	result := download.GetHtml(URL, constants.PATH_FILE+constants.HTML_NAME)
	if result {
		fmt.Println("file is downloaded successfully.")
	}

	log.Println("end: assignment", ASSIGNMENT_NUMBER)

}
