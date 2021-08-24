package main

import "fmt"

func main() {
	//1.
	s := "i'm sorry dave i can't do that"
	fmt.Println(s)

	//2.
	convertToByte := []byte(s)
	fmt.Println(convertToByte)

	//3.
	convertToString := string(convertToByte)
	fmt.Println(convertToString)

	//4.
	slicedString1 := s[:14]
	fmt.Println(slicedString1)

	//5.
	slicedString2 := s[10:22]
	fmt.Println(slicedString2)

	//6.
	slicedString3 := s[17:]
	fmt.Println(slicedString3)

	//7.
	for _, v := range []byte(s) {
		fmt.Println(string(v))
	}

}
