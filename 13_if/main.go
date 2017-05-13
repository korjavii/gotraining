package main

import (
	"fmt"
)

func main() {
	b := false

	if food := "bacon"; b {
		fmt.Println(food)
	} else {
		fmt.Println(b)
	}
}
