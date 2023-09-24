package main

import (
	"fmt"
	"net/http"

	"github.com/anik36/Golang/tree/advanced-go/constants"
	"github.com/anik36/Golang/tree/advanced-go/handlers"
)

func main() {

	http.HandleFunc("/home", handlers.Home)
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", constants.PORT_NUMBER)
	_ = http.ListenAndServeTLS(constants.PORT_NUMBER, constants.SERVER_CERT, constants.SERVER_KEY, nil)
}
