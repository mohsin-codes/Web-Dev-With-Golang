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

	//Used casting to cast g1 as int
	x = int(g1)
	fmt.Println(x)
}
