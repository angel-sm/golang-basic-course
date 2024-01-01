package main

import "fmt"

func main() {
	switch module := 6 % 2; module {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condiciones
	valie := 1
	switch {
	case valie > 100:
		fmt.Println("Mayor a 100")
	case valie < 0:
		fmt.Println("Es menor a 0")
	default:
		fmt.Println("No condicion")
	}
}
