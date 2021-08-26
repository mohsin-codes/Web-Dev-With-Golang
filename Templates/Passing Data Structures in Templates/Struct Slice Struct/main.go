package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type people struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type items struct {
	PeopleMentioned []people
	Transport       []car
}

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

	f := car{
		Manufacturer: "Ford",
		Model:        "F-150",
		Doors:        2,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	importantPeople := []people{buddha, gandhi, aL}
	cars := []car{f, c}

	data := items{
		PeopleMentioned: importantPeople,
		Transport:       cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
