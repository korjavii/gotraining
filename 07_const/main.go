package main

import "fmt"

const p string = "cowboy"
const (
	PI   = 3.14
	TEST = iota
)

func main() {
	const a = 10

	fmt.Println("p - ", p)
	fmt.Println("a - ", a)
	fmt.Println("pi - ", PI)
	fmt.Println("test - ", TEST)
}
