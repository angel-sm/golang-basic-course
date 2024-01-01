package main

import "fmt"

func main() {
	nombre := "Jose Angel"
	lastName := "Pacheco"

	// Println
	fmt.Println("Hello world")

	// Printf
	fmt.Printf("%s %s te manda\n salu2", nombre, lastName)
	// %s String
	// %b Int
	// if you dot know type of varuable you can add %v

	// Sprintf
	message := fmt.Sprintf("\n%s %s te manda\n salu2", nombre, lastName)
	fmt.Println(message)

	// Variable type
	fmt.Printf("Type of value: %T", nombre)
}
