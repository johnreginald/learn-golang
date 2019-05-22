package main

import "fmt"

func main() {
	// Array, Slice, Map are Reference Type in Go
	test := []int{10, 20, 30}
	change(test)
	fmt.Println(test)
}

func change(data []int) {
	fmt.Println(data)
	data[0] = 50
}
