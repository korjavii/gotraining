package main

import "fmt"

var g string = "cowboy"

func main() {
	a := 10
	b := "golang"
	c := 3.14
	d := true

	fmt.Printf("%v - type: %T \n", a, a)
	fmt.Printf("%v - type: %T \n", b, b)
	fmt.Printf("%v - type: %T \n", c, c)
	fmt.Printf("%v - type: %T \n", d, d)
	fmt.Printf("%v - type: %T \n", g, g)
}
