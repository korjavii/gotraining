package main

import "fmt"

var x int = 42
var y string = "go"

func main() {
	fmt.Println(x)
	foo()

	z := 43
	fmt.Println(z)
}

func foo() {
	fmt.Println(y)
}
