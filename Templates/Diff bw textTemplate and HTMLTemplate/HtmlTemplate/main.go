package main

import (
	"html/template"
	"log"
	"os"
)

type items struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main() {
	v1 := items{
		Title:   "Heading",
		Heading: "This is the Heading",
		Input:   `<script>alert("Hi there!");</script>`,
	}

	err := tpl.Execute(os.Stdout, v1)
	if err != nil {
		log.Fatalln(err)
	}
}
