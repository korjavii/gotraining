package main

import "fmt"

func main() {
	fmt.Println(greet("Test", "222"))
	fmt.Println(average(1, 2, 4, 5.6))
	data := []float64{2, 5, 9}
	fmt.Println(average(data...))

	visit([]int{1, 2, 3}, func(arg2 int) {
		fmt.Println(arg2)
	})
}

func greet(name, name2 string) (string, string) {
	return fmt.Sprint(name, name2), fmt.Sprint(name, name2)
}

func average(sf ...float64) float64 {
	total := 0.0

	for _, v := range sf {
		total += v
	}

	return total / float64(len(sf))
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
