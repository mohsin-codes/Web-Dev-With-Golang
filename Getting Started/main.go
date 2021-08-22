package main

import "fmt"

//struct declarartion
type person struct {
	fname string
	lname string
}

type secretagent struct {
	person
	licenceToKill bool
}

//Interface
type human interface {
	speak()
}

//Function Declaration
//using Reciever Function
func (p person) speak() {
	fmt.Println(p.fname + ` says, "Hello There"`)
}

//embedded function declaration
func (sa secretagent) speak() {
	fmt.Println(sa.person.fname + sa.lname + ` says, "Shaken, not stirred"`)
}

//Interface function declaration
func saySomething(h human) {
	h.speak()
}

//global variable declaraion
var temp int

func main() {
	//Short Variable Declaration
	x := 7
	fmt.Println(x)

	//Printing type of variable
	fmt.Printf("%T\n", x)

	fmt.Println(temp)

	//Data Structures
	//Using composite literal - []type{data}
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums)

	//Maps - using composite literals
	m := map[string]int{
		"Mohsin": 22,
		"Ashar":  17,
	}
	fmt.Println(m)

	//Struct Initialization
	p1 := person{
		"Mojo",
		"Rojo",
	}
	fmt.Println(p1)

	james := secretagent{
		person{
			"James",
			"Bond",
		},
		true,
	}

	//Function call
	p1.speak()
	james.speak()

	//Function calls through interface
	saySomething(p1)
	saySomething(james)
}
