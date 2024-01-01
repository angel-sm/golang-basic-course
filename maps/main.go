package main

import "fmt"

func main() {
	m := make(map[string]int)
	mv2 := map[string]int{
		"Luis": 10,
	}

	fmt.Println(mv2)

	m["Angel"] = 25
	m["Pedro"] = 17

	// iterate map

	for i, v := range m {
		fmt.Println(i, v)
	}

	// get any value

	value, ok := m["Angels"]
	fmt.Println(value, ok)
}
