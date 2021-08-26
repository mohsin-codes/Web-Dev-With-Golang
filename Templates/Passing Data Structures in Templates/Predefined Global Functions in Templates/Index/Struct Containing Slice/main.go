package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseFiles("1.gohtml"))
}

func main(){
	xs := []string{"one", "two", "three", "four", "five"}

	data := struct{
		Words []string
		Lname string
	}{
		xs,
		"Mississippi",
	}

	err := tpl.Execute(os.Stdout, data)
	if err!=nil{
		log.Fatalln(err)
	}
}