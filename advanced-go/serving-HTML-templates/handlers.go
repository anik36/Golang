package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("/media/aniket/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Documents/GIT ANIKET/advanced-go/templates/" + tmpl)
	if err != nil {
		log.Fatal("Unable to parse from template:", err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
