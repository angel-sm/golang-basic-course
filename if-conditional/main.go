package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	value1 := 10
	value2 := 20

	if value1 == 1 {
		fmt.Println("Es 1")
	} else {
		fmt.Println("No es 1")
	}

	// AND
	if value1 == 10 && value2 == 20 {
		fmt.Println("Es verdad AND")
	}

	// OR
	if value1 == 10 || value2 == 40 {
		fmt.Println("Es verdad OR")
	}

	//
	value, err := strconv.Atoi("54")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value %d", value)

}
