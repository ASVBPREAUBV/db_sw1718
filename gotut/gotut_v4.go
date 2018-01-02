package main

import "fmt"

func main() {
	x := 15
	a := &x //actual memory address
	fmt.Println("a", a);
	fmt.Println("*a", *a);
	*a = 5
	fmt.Println("x", x);
	*a = *a * *a
	fmt.Println("x", x)
	fmt.Println("*a", *a)
}