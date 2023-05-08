package main

import (
	"fmt"
	"net/http"

	"github.com/adityasinghtads/go-web-app/pkg/handlers"
)

const portNumber = ":8080"

// the main application function.
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting Application at %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}

// Function can be called in the package, Upper case can be used outside of pacakge or incase lower case is not accesible
