package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) // execute the first template matched
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.abc", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.whatever_name", nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "end.dnsakjdnals", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
