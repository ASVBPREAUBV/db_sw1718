package main

import (
	"fmt"
)

type car struct {
	gas_pedal   uint16 //min 0 max 65535
	break_pedal uint16 //min 0 max 65535
	name        string
	speed       float64
}

func main() {
	var car1_gaspedal uint16 = 0
	car1 := car{car1_gaspedal, 313, "Peter", 200}
	fmt.Println("car1", car1)
	fmt.Println("car1.name", car1.name)

}

