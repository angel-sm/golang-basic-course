package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func returnDobleValue(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola mundo")
	tripleArgument(1, 2, "Hola")

	// Return value
	result := returnValue(2)
	fmt.Println(result)

	// Get all values
	v1, v2 := returnDobleValue(3)
	fmt.Println(v1, v2)

	// Scape value
	res, _ := returnDobleValue(4)
	fmt.Println(res)
}
