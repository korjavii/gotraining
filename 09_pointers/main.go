package main

import (
	"fmt"
)

func main() {
	var a float64
	a = 42

	fmt.Println(a)
	fmt.Println(&a)

	var b *float64 = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 43
	fmt.Println(a)

	x := 1
	fmt.Printf("%p \n", &x)
	fmt.Println(&x)
	zero(&x)
	fmt.Println(x)
}

func zero(x *int) {
	fmt.Printf("%p \n", &x)
	fmt.Println(x)
	*x = 0
}
