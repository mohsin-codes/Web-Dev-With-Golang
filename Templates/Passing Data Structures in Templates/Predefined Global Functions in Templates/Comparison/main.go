package main

import (
	"log"
	"os"
	"text/template"
)

type numbers struct {
	Num1 int
	Num2 int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	s1 := numbers{
		10,
		20,
	}

	err := tpl.Execute(os.Stdout, s1)
	if err != nil {
		log.Fatalln(err)
	}
}
