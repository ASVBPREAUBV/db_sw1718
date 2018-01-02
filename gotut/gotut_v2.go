package main

import (
	"fmt"
	"math"
	"math/rand"
)

func foo() {
	fmt.Println(math.Log(12), "Hello Go")
}

func bar() {
	fmt.Println(rand.Intn(100), "rand")
}

func main() {
	foo()
	bar()
}