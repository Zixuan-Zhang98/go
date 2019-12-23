package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func main() {
	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/index.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	tpl.Execute(w, nil)
}
