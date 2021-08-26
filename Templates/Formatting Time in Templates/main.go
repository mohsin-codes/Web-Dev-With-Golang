package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("1.gohtml"))
}

func monthDayYear(t time.Time) string {
	return t.Format("02-01-2006")
}

var fm = template.FuncMap{
	"fDateDMY": monthDayYear,
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "1.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
