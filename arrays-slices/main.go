package main

import "fmt"

func main() {
	// Array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	fmt.Println(arr, len(arr), cap(arr))

	// Slice
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice, len(slice), cap(slice))

	//Slice methods
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[2:])

	// [(First inclusive):(Second exclusive)]

	// Append
	slice = append(slice, 10)
	fmt.Println(slice)

	// Append list
	newSlice := []int{11, 12, 13}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
