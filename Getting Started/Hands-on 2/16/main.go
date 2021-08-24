package main

import "fmt"

type gator int

type flamingo bool

type swampCreature interface{
	greeting() string
}

//Adding a method to type gator
func (g gator) greeting() string {
	return "Hello, I am a gator."
}

func (f flamingo) greeting() string{
	return "Hello, I am pink and beautiful and wonderful."
}

func bayou(sc swampCreature){
	fmt.Println(sc.greeting())
}

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

	fmt.Println(g1.greeting())

	//Creating flamingo identifier
	var f1 flamingo
	f1 = true
	fmt.Println(f1)


	//Calling interface functions
	bayou(g1)
	bayou(f1)
}
