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
	inspire := []string{"Gandhi", "APJ Abdul Kalam", "Elon Musk", "Albert Einstein", "Isaac Newton"}

	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", inspire)
	if err != nil {
		log.Fatalln(err)
	}
}
