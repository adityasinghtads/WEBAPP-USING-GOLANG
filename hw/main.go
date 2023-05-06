package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}

// About page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is about page and 2+2 is %d", sum))
	//fmt.Fprintf(w, "This is about page")
}

// AddValues function, takes 2 Int and returns sum.
func addValues(x, y int) int {
	return x + y
}

// the main application function.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting Application at %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

// Function can be called in the package, Upper case can be used outside of pacakge or incase lower case is not accesible
