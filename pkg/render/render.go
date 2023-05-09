package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error in parsed template ", err)
	}
	//fmt.Println("Hey hello")
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	// Using Map. for storing data of the template, so that each call is not made from the disk.
	var tmpl *template.Template
	var err error
	// CHeck to see if we already have the template in cache
	_, inMap := tc[t]
	if !inMap {
		log.Println("Creating template and adding to cache.")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("Using Cache template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	//using map for storing data.
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
	}
	//psarse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache
	tc[t] = tmpl
	return nil
}
