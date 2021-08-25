package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("temp.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("temp2.gohtml")
	if err != nil {
		log.Println("Error in Creating file")
	}
	defer file.Close()

	err = tpl.Execute(file, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
cd 