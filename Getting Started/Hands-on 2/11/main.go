package main

import "fmt"

type gator int

func main() {
	var g1 gator
	g1 = 19
	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
}
