package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", 42)
	if err != nil {
		log.Println(err)
	}
}
