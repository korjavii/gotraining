package main

import "fmt"

var x int = 42
var y string = "go"

func main() {
	fmt.Println(x)
	foo()

	z := 43
	fmt.Println(z)
	{
		zz := 33
		fmt.Println(zz)
	}
	fmt.Println(increment())
}

func foo() {
	fmt.Println(y)
}

func increment() int {
	x++
	return x
}
