package main

import (
	"fmt"
	"log"

	"redmine.merce.co/git/go-learning-project/anikets/constants"
	"redmine.merce.co/git/go-learning-project/anikets/fileCompress26/fileconv"
	"redmine.merce.co/git/go-learning-project/anikets/fileHtmlDownload27/download"
)

/******************************************************************************************
* 26. Visit https://merce.co and save the homepage as an HTML file from a browser as
* “merce-homepage.html”. Write a program to read the saved HTML file, compress it and store
* the compressed file as “merce-homepage.html.zip”. Use ready classes for compression. At
* the end, the program shall print HTML file size, Compressed file size. [file operation
* concepts, 8-12 hours]
*******************************************************************************************/

const (
	ASSIGNMENT_NUMBER = 26
	URL               = "https://merce.co"
)

func main() {
	log.Println("started: assignment", ASSIGNMENT_NUMBER)

	result := download.GetHtml(URL, constants.PATH_FILE+constants.HTML_NAME)
	if result {
		fmt.Println("file is downloaded successfully.")
	}

	isErr := fileconv.File2Zip(constants.PATH_FILE, constants.HTML_NAME, constants.ZIP_NAME)
	if !isErr {
		log.Fatal(constants.ERROR_26_5)
	}

	fmt.Println("HTML file size:", fileconv.FileSize(constants.PATH_FILE+constants.HTML_NAME), "Bytes")
	fmt.Println("Compressed file size:", fileconv.FileSize(constants.PATH_FILE+constants.ZIP_NAME), "Bytes")

	log.Println("end: assignment", ASSIGNMENT_NUMBER)
}
