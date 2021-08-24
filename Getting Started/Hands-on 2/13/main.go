package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 27
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	//Cannot be used since g1 is of type gator and x is of type int
	x = g1
	fmt.Println(x)
	fmt.Println(g1)
}
