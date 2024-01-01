package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (PC pc) Ping() {
	fmt.Println(PC.brand, "Pong")
}

func (PC *pc) DuplicateRam() {
	PC.ram = PC.ram * 2
}

func (Pc *pc) PrintPC() {
	fmt.Println(Pc)
}

func main() {

	myPC := pc{
		ram:   8,
		disk:  200,
		brand: "Asus",
	}

	myPC.Ping()

	myPC.PrintPC()

	myPC.DuplicateRam()
	myPC.DuplicateRam()

	myPC.PrintPC()

}
