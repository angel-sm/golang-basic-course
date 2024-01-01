package main

import "fmt"

func conditionFor() {
	fmt.Println("Start condition for")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func forWhile() {
	fmt.Println("Start for while")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
}

func forForever() {
	counter := 0
	for {
		fmt.Println(counter)
		counter++
	}
}

func main() {
	conditionFor()
	forWhile()

}
