package main

import "fmt"

func main() {
	x, y := test()
	fmt.Println(x, y)
}

func test() (int, int) {
	a := 50
	b := 60
	return a, b
}
