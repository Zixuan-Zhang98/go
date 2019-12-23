package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("templates/index.gohtml"))
	tpl.Execute(w, nil)
}
