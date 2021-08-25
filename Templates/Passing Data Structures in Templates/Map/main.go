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
	importantPeople := map[string]string{
		"India":          "Gandhi",
		"Germany":        "Albert Einstein",
		"United Kingdom": "Isaac Newton",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", importantPeople)
	if err != nil {
		log.Fatalln(err)
	}
}
