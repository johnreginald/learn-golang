package main

import "fmt"

func main() {
	for _, value := range []int{10, 20, 30} {
		defer fmt.Println(value)
	}
}
