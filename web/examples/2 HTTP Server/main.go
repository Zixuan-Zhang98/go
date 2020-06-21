package main

import (
	"fmt"
	"net/http"
)

/*
A basic HTTP server has a few key jobs to take care of:
	1. Process dynamic requests: Process incoming requests from users who browse the website, log into their accounts or post images.
	2. Serve static assets: Serve JavaScript, CSS and images to browsers to create a dynamic experience for the user.
	3. Accept connections: The HTTP Server must listen on a specific port to be able to accept connections from the internet.

1. Process dynamic requests:
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!")
	})

2. Serving static assets:
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

3. Accept connections:
	http.ListenAndServe(":80", nil)
*/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
