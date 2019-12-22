package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	// tpl is a pointer to a Template
	// Template is just a container which contains all the parsed files
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

//$ go run main.go
//<!DOCTYPE html>
//<html lang="en">
//<head>
//    <meta charset="UTF-8">
//    <title>Hello World!</title>
//</head>
//<body>
//<h1>Hello</h1>
//</body>
//</html>