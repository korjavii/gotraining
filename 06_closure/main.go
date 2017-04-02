package main

import "fmt"

func main() {
	z := 43
	increment := func() int {
		z++
		return z
	}
	fmt.Println(increment())

	incr := wrapper()
	fmt.Println(incr())
	fmt.Println(incr())
}

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
