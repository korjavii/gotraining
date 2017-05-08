package main

import (
	"fmt"
)

func main() {
	switch "b" {
	case "a":
		fmt.Println("a")
	case "b", "d":
		fmt.Println("b")
		fallthrough
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("default")
	}
}
