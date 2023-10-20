package renders

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

var functions = template.FuncMap{}

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

func CreateTemplate() error {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseFiles("./templates/*.layout.tmpl")
			if err != nil {
				return err
			}
		}
		myCache[name] = templateSet
	}
}
