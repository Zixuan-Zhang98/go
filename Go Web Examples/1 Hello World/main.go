package main

import (
	"fmt"
	"net/http"
)

/*
The net/http package from the standard library contains all functionality about the HTTP protocol.
This includes (among many other things) an HTTP client and an HTTP server.

1. Registering a Request Handler:
	A handler in Go is a function with this signature:
		func (w http.ResponseWriter, r *http.Request)

	The function receives two parameters:
	An http.ResponseWriter which is where you write your text/html response to.
	An http.Request which contains all information about this HTTP request including things like the URL or header fields.

2. Registering a request handler to the default HTTP Server is as simple as this:
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

3. Listen for HTTP Connections:
	The request handler alone can not accept any HTTP connections from the outside.
	An HTTP server has to listen on a port to pass connections on to the request handler.
		e.g. http.ListenAndServe(":80", nil)
*/

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprintf(writer, "Hello, you've requested: %s\n", request.URL.Path)
		if err != nil {
			fmt.Println(err)
		}
	})

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println(err)
	}
}
