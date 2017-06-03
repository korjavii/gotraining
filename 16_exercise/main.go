package main

import "fmt"

func main() {
	fmt.Println(half(2, func(arg2 int) bool {
		return arg2%2 == 0
	}))

	data := []int{2, 5, 9, 5567899}
	fmt.Println(greatest(data...))

	fmt.Println(((true && false) || (false && true) || !(false && false)))
}

func half(n int, callback func(int) bool) (int, bool) {
	h := n / 2
	b := callback(n)
	return h, b
}

func greatest(sf ...int) int {
	gr := 0

	for _, v := range sf {
		if v > gr {
			gr = v
		}
	}

	return gr
}
