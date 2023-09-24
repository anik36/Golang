package renders

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("/media/aniket/9c8cbcfa-7a63-4b30-910a-2586f19beb53/Aniket/Documents/GIT ANIKET/advanced-go/templates/" + tmpl)
	if err != nil {
		log.Fatal("Unable to parse from template:", err)
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
