package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	optimusPrime := truck{
		vehicle{
			2,
			"Blue",
		},
		true,
	}

	bumblebee := sedan{
		vehicle{
			4,
			"Yellow",
		},
		true,
	}

	fmt.Println(optimusPrime)
	fmt.Println(bumblebee)

	fmt.Println(optimusPrime.vehicle)
	fmt.Println(bumblebee.vehicle)
}
