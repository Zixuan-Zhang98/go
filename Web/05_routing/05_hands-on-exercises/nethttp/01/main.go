package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "index")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "doggy")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Zixuan")
}
