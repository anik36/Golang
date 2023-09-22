package main

import (
	"fmt"
	"log"
	"net/http"

	applicationproperties "github.com/anik36/Golang/tree/advanced-go/application-properties"
)

func main() {

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		totalByte, err := fmt.Fprintf(responseWriter, "Hello world!")
		fmt.Printf("Number of bytes written: %d \n", totalByte)
		if err != nil {
			log.Println(err)
		}
	})

	_ = http.ListenAndServe(applicationproperties.PORT_NUMBER, nil)
}
