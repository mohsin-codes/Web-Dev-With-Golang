package main

import (
	"log"
	"os"
	"text/template"
)

type people struct {
	Name  string
	Motto string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	buddha := people{
		Name:  "Buddha ",
		Motto: "Eat Burgers",
	}

	gandhi := people{
		Name:  "MK Gandhi",
		Motto: "Eat Pizza",
	}

	aL := people{
		Name:  "Abraham Lincoln",
		Motto: "Take a chill pill",
	}

	importantPeople := []people{buddha, gandhi, aL}
	err := tpl.Execute(os.Stdout, importantPeople)
	if err != nil {
		log.Fatalln(err)
	}
}
