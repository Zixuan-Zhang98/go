package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `
	<!-- not serving file on our server -->
	<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/toby.jpg">
	`)
}
