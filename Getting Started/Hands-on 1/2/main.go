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

func (p person) pSpeak() string {
	// fmt.Println(p.fname, `says, "Hello There!"`)
	return (p.fname + ` says, "Hello There!"`)
}

func (sa secretAgent) saSpeak() string {
	// fmt.Println(sa.fname, `says, "Shaken, not stirred"`)
	return (sa.fname + ` says, "Shaken, not stirred"`)
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
	fmt.Println(p1.pSpeak())
	fmt.Println(Jesse.saSpeak())
}
