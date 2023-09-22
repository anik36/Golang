package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	PORT_NUMBER = ":9094"
	SERVER_CERT = "/media/aniket/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Documents/GIT ANIKET/advanced-go/SSL/server.crt"
	SERVER_KEY  = "/media/aniket/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Documents/GIT ANIKET/advanced-go/SSL/server.key"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	div, err := divideValue(2.0, 0.0)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintln(w, "This is the about page and 2 / 0 is :", div)
}

// sum adds two integer and return the sum
func sum(a, b int) int {
	return a + b
}

func divideValue(a, b float32) (float32, error) {
	if b == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	return a / b, nil
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
	http.HandleFunc("/divide", Divide)

	fmt.Println("Starting application on port", PORT_NUMBER)
	_ = http.ListenAndServeTLS(PORT_NUMBER, SERVER_CERT, SERVER_KEY, nil)
}
