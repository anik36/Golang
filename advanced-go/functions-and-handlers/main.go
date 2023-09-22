package main

import (
	"fmt"
	"log"
	"net/http"

	applicationproperties "github.com/anik36/Golang/tree/advanced-go/application-properties"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := sum(2, 2)
	fmt.Fprintln(w, "This is the about page and 2 + 2 is :", sum)
}

// sum adds two integer and return the sum
func sum(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		totalByte, err := fmt.Fprintf(responseWriter, "Hello world!")
		fmt.Printf("Number of bytes written: %d \n", totalByte)
		if err != nil {
			log.Println(err)
		}
	})

	http.HandleFunc("/home", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port", applicationproperties.PORT_NUMBER)
	_ = http.ListenAndServeTLS(applicationproperties.PORT_NUMBER, applicationproperties.SERVER_CERT, applicationproperties.SERVER_KEY, nil)
}
