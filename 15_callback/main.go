package main

import "fmt"

func main() {
	defer fmt.Println(1111)
	xs := filter([]int{1, 5, 3}, func(arg2 int) bool {
		return arg2 > 1
	})

	fmt.Println(xs)

	abc := make([]string, 1, 25)
	fmt.Println(abc)
	passByRef(abc)
	fmt.Println(abc)
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func passByRef(str []string) {
	str[0] = "test"
}
