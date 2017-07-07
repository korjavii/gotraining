package main

import "fmt"

func main() {
	x := []string{"a", "b", "c", "d", "g", "z"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[2:4])
	fmt.Println(x[2])
	fmt.Println("mySlice"[2])

	mySlice := make([]int, 0, 5)
	mySlice = append(mySlice, 1)
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	myOtherSlice := []int{1, 2, 3, 4, 5}
	mySlice = append(mySlice, myOtherSlice...)
	fmt.Println(mySlice)
	mySlice = append(mySlice[:2], myOtherSlice[3:]...)
	fmt.Println(mySlice)

	records := make([][]string, 0)
	row := make([]string, 2)
	row[0] = "ivan"
	row[1] = "ivanov"

	records = append(records, row)
	fmt.Println(records)
}
