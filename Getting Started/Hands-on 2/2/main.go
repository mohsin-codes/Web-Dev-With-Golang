package main

import "fmt"

func main()  {
	createdMap := map[string]int{
		"Peter" : 1,
		"Mary Jane" : 2,
		"Doctor Strange" : 3,
	}

	fmt.Println(createdMap)

	for i := range createdMap{
		fmt.Println(i)
	}

	for i, val := range createdMap{
		fmt.Println(i, val)
	}
}