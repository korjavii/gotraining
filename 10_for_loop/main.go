package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, "-", j)
		}
	}

	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		if i > 10 {
			break
		}
	}

	for i := 50; i <= 140; i++ {
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}
}
