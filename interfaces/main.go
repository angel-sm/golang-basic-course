package main

import "fmt"

type figures2D interface {
	area() float64
}

type square struct {
	base float64
}

func (Square square) area() float64 {
	return Square.base * Square.base
}

type rectangle struct {
	base   float64
	heigth float64
}

func (Rectangle rectangle) area() float64 {
	return Rectangle.base * Rectangle.heigth
}

func calculate(f figures2D) {
	fmt.Println("Area", f.area())
}

func main() {
	mySqueare := square{
		base: 10,
	}

	myRectangle := rectangle{
		base:   10,
		heigth: 20,
	}

	calculate(myRectangle)
	calculate(mySqueare)

	// interface list

	myInterfaceList := []interface{}{
		2, "Hola", true, 1.3,
	}

	fmt.Println(myInterfaceList...)

}
