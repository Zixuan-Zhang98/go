package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.abc")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	tpl, err = tpl.ParseFiles("two.whatever_name", "end.dnsakjdnals")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "end.dnsakjdnals", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.whatever_name", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.abc", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) // Execute() method will only execute the first template in all templates
	if err != nil {
		log.Fatalln(err)
	}
}
