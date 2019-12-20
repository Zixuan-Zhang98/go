package main

import "fmt"

func main() {
	name := "Zixuan Zhang"
	// string concatenation as templating
	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`
	fmt.Println(tpl)
	// use "go run main.go > index.html" to import what's printed out into a file
	// html are just text
}
