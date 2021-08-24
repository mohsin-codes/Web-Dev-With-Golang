package main

import "fmt"

type person struct {
	fname string
	lname string
}

func main() {
	p1 := person{
		"Super",
		"Mario Jr.",
	}

	fmt.Println(p1)

	fmt.Println(p1.fname, p1.lname)
}
