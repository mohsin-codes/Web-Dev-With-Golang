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

func (t truck) transportationDevice() string {
	return "A truck or lorry is a motor vehicle designed to transport cargo, carry specialized payloads, or perform other utilitarian work"
}

func (s sedan) transportationDevice() string {
	return "A sedan, or saloon is a passenger car in a three-box configuration with separate compartments for engine, passenger, and cargo"
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

	fmt.Println(optimusPrime.transportationDevice())
	fmt.Println(bumblebee.transportationDevice())
}
