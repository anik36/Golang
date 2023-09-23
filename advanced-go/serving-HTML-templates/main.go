package main

import (
	"fmt"
	"net/http"

	"github.com/anik36/Golang/tree/advanced-go/constants"
)

func main() {

	http.HandleFunc("/home", Home)
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Starting application on port", constants.PORT_NUMBER)
	_ = http.ListenAndServeTLS(constants.PORT_NUMBER, constants.SERVER_CERT, constants.SERVER_KEY, nil)
}
