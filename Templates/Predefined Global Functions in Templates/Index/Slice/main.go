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
	cities := []string{"Copanhagen", "New Delhi", "Sydney", "Toronto", "Wellington"}

	err := tpl.Execute(os.Stdout, cities)
	if err != nil {
		log.Fatalln(err)
	}
}
