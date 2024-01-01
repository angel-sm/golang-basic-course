package main

import "fmt"

func main() {
	const base = 10
	const height = 10
	area := base * base

	fmt.Println(area)

	x := 10
	y := 20

	// addition
	addition := x + y
	fmt.Println("addition:", addition)

	// subtract
	subtract := x - y
	fmt.Println("subtract:", subtract)

	// multiplication
	multiplication := x * y
	fmt.Println("multiplication: ", multiplication)

	// division
	division := y / x
	fmt.Println("division: ", division)

	// Modulo
	modulo := y % x
	fmt.Println("modulo: ", modulo)

	// increment
	x++
	fmt.Println("increment: ", x)

	// decrement
	x--
	fmt.Println("decrement: ", x)

	// Challange
	// Rectangle area
	length := 10
	width := 20
	rectangleArea := length * width
	fmt.Println("Rectangle area is:", rectangleArea)

	// Circle area
	pi := 3.1416
	radio := 20

	circleArea := pi * (float64(radio * radio))
	fmt.Println("Circle area is:", circleArea)
}
