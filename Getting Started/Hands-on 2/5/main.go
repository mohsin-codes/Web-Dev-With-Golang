package main

import "fmt"

type person struct {
	fname string
	lname string
	favFood []string
}

func main() {
	p1 := person{
		"Super",
		"Mario Jr.",
		[]string{"Mushroom", "Flowers", "Stars", "Coins"},
	}

	fmt.Println(p1.favFood)
	for _, val := range p1.favFood{
		fmt.Println(val)
	}
}
