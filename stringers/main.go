package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (Pc pc) String() string {
	return fmt.Sprintf("Tengo un %d GB RAM, %d GB DD, y es una %s", Pc.ram, Pc.disk, Pc.brand)
}

func main() {
	myPc := pc{
		ram:   8,
		disk:  100,
		brand: "ASUS",
	}

	fmt.Println(myPc.String())
}
