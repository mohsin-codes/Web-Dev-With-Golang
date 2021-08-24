package main

import "fmt"

func main() {
	createdSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(createdSlice)

	for i := range createdSlice {
		fmt.Println(i)
	}

	for i, val := range createdSlice {
		fmt.Println(i, val)
	}
}
