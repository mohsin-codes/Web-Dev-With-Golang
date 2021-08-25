package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type importantPeople struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	buddha := importantPeople{
		Name:  "Buddha",
		Motto: "Eat Burgers",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
