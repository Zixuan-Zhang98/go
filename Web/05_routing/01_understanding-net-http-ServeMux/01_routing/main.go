package main

import (
	"io"
	"net/http"
)

type Rebecca int

func (e Rebecca) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "doggy")
	case "/cat":
		io.WriteString(w, "kitty")
	}
}

func main() {
	var e Rebecca
	http.ListenAndServe(":8080", e)
}
