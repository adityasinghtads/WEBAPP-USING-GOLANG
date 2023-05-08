package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	a := 100.0
	b := 0.0
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", a, b, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by 0")
		return 0, err
	}
	result := x / y
	return result, nil
}

// the main application function.
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting Application at %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

// Function can be called in the package, Upper case can be used outside of pacakge or incase lower case is not accesible
