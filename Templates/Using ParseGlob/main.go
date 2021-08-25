package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseGlob("files/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "1.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
	
	err = tpl.ExecuteTemplate(os.Stdout, "3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
