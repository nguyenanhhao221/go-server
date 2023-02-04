package main

import (
	"fmt"
	"log"
	"net/http"
)

// In Go, http.Response is a struct that represents an HTTP response and http.ResponseWriter is an interface that defines the methods that a response writer must have.
// http.ResponseWriter is often used in HTTP handler functions to construct the HTTP response, while http.Response is a struct that represents the actual HTTP response sent to the client.
// When a HTTP handler function is called, it receives an http.ResponseWriter as an argument. The handler then uses the methods provided by the http.ResponseWriter interface to write the HTTP response, such as the status code, headers, and body.
// The advantage of using http.ResponseWriter is that it provides a standard interface for writing HTTP responses, which allows different types of response writers to be used interchangeably. For example, you can use a custom implementation of http.ResponseWriter to capture and inspect the HTTP response before it is sent to the client.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not accepted", http.StatusMethodNotAllowed)
		return
	}
	// fmt.Printf is used to print data to the terminal, but when you're building a web server with Go, you want to send the data to the user's browser, not just the terminal. That's why you use fmt.Fprintf instead. It sends the data to the user's browser instead of the terminal, which is what you need when you're building a web server.
	fmt.Fprintf(w, "Hello World")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != "POST" {
		http.Error(w, "method is not accepted", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "POST Request success")
}

func main() {
	// Create a `FileServer` that serves files from the "./static" directory
	fileServer := http.FileServer(http.Dir("./static"))

	// Use the `Handle` function to bind the `fileServer` to the root URL ("/")
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	// Print a message indicating that the server is starting
	fmt.Printf("Starting server at port 8080\n")

	// Start the HTTP server using `ListenAndServe` on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		// If there was an error starting the server, log the error and exit
		log.Fatal(err)
	}

}
