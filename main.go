package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = ":8080"

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemp, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemp.Execute(w, nil)
	if err != nil {
		fmt.Println("error in parsed template ", err)
	}
}

// the main application function.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Application at %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

// Function can be called in the package, Upper case can be used outside of pacakge or incase lower case is not accesible
