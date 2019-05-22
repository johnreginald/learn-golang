package main

import "fmt"

func mainn() {
	x := 50
	add(&x)
	fmt.Println(x)
}

func add(x *int) int {
	*x = *x + 1
	return *x
}
