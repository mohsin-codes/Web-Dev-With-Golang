package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	//Parsing the file using template.ParseFiles
	tpl, err := template.ParseFiles("1.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//Printing out the read template as the only available template yet
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Parsing multiple files using template.ParseFiles
	tpl, err = tpl.ParseFiles("2.gohtml", "3.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//Printing out the exact template as mentioned in the ExecuteTemplate method
	err = tpl.ExecuteTemplate(os.Stdout, "2.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Printing out the exact template as mentioned in the ExecuteTemplate method
	err = tpl.ExecuteTemplate(os.Stdout, "3.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	//Printing out all the templates again will print only the topmost template ever recieved
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
