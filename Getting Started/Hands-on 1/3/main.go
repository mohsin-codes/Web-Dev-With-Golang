package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak() string
}

func (p person) speak() string {
	// fmt.Println(p.fname, `says, "Hello There!"`)
	return (p.fname + ` says, "Hello There!"`)
}

func (sa secretAgent) speak() string {
	// fmt.Println(sa.fname, `says, "Shaken, not stirred"`)
	return (sa.fname + ` says, "Shaken, not stirred"`)
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p1 := person{
		"Kill",
		"Monger",
	}

	Jesse := secretAgent{
		person{
			"Jesse",
			"Team Rocket",
		},
		true,
	}

	fmt.Println(p1.fname)
	fmt.Println(Jesse.fname)
	fmt.Println(p1.speak())
	fmt.Println(Jesse.speak())

	saySomething(p1)
	saySomething(Jesse)
}
