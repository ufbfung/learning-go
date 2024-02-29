package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a struct to decode the JSON data into
type GreetingRequest struct {
	Name string `json:"name"`
}

func main() {
	// Handle root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Handle /greet path for GET
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			name := r.URL.Query().Get("name")
			if name == "" {
				name = "Guest"
			}
			fmt.Fprintf(w, "Hello, %s!", name)
		} else if r.Method == "POST" {
			// Handle POST request
			var greetReq GreetingRequest
			// Decode the JSON body into the GreetingRequest struct
			err := json.NewDecoder(r.Body).Decode(&greetReq)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Fprintf(w, "Hello, %s!", greetReq.Name)
		} else {
			// If the method is not supported, return a 405 Method Not Allowed
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
