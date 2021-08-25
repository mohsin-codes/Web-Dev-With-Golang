package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", "Where there is a will, there is a way")
	if err != nil {
		log.Fatalln(err)
	}
}
