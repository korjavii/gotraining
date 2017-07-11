package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 11
	fmt.Println(m)
	fmt.Println(len(m))
	delete(m, "a")
	fmt.Println(m)
	v, ok := m["b"]
	fmt.Println("ok?", ok, v)
	test := greeting()
	fmt.Println(test)
	if val, exists := test["Hi"]; exists {
		fmt.Println(val)
	}
}

func greeting() map[string]string {
	myGreeting := map[string]string{
		"Hello ": "myke",
		"Hi":     "john",
	}
	return myGreeting
}
