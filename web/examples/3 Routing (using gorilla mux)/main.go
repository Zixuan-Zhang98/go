package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
1. Installing the gorilla/mux package:
	go get -u github.com/gorilla/mux

2. Creating a new Router:
	r := mux.NewRouter()

3. Registering a Request Handler:
	r.HandleFunc(...)

4. URL Parameters:
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		// get the book
		// navigate to the page
	})

	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		vars["title"] // the book title slug
		vars["page"] // the page
	}

5. Setting the HTTP server’s router:
	http.ListenAndServe(":80", r)
*/

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	http.ListenAndServe(":80", r)
}
