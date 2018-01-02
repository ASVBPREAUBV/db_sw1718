package main

import "fmt"

func test(x string, y string) string {
	return x + "  " + y
}

func main() {
	var string1 string = "sdfg"
	fmt.Println("print test", test(string1, "dfg"));
}