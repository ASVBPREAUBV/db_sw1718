package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("Defer 1")
	}()
	defer func() {
		fmt.Println("Defer 2")
	}()

	fmt.Println("I'm not deferred!")
}