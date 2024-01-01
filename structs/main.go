package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{
		brand: "Ford",
		year:  2023,
	}

	fmt.Println(myCar)

	var secondCar car
	secondCar.brand = "Chevrolet"

	fmt.Println(secondCar)
}
