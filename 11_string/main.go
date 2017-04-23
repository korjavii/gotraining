package main

import (
	"fmt"
)

func main() {
	foo := 'a'
	fmt.Println(foo, "abcd")
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", rune(foo))
}
