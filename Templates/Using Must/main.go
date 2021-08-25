package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("Files/*.gohtml"))
}

func main() {
	err := tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "1.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "2.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "3.gohtml", nil)
	if err != nil {
		log.Println(err)
	}

}
