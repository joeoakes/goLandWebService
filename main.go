package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Define the route handler function
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	// Register the route handler function to handle requests for the "/hello" path
	http.HandleFunc("/hello", helloHandler)

	// Start the web server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}
