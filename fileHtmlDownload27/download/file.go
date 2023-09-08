package download

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
)

func GetHtml(url string, path string) bool {
	// ignore ssl certification
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: transport}
	// Send an HTTP GET request to the web page
	response, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// Convert data to byte
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// write whole the body byte to file
	err = os.WriteFile(path, body, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return true
}
