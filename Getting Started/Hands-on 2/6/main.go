package main

import "fmt"

type person struct {
	fname   string
	lname   string
	favFood []string
}

func (p person) walk() string {
	return (p.fname + ` ` + p.lname + ` is walking`)
}

func main() {
	p1 := person{
		"Super",
		"Mario Jr.",
		[]string{"Mushroom", "Flowers", "Stars", "Coins"},
	}

	s := fmt.Sprintln(p1.walk())
	fmt.Println(s)
}
