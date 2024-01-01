package main

import "fmt"

func isPalindrome(text string) bool {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	return textReverse == text
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, v := range slice {
		fmt.Println("Indice", i)
		fmt.Println("Value", v)
	}

	// ama

}
