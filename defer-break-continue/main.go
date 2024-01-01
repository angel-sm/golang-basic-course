package main

import "fmt"

func main() {

	// Defer
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue and break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 4 {
			fmt.Println("Es 4")
			break
		}
	}
}
