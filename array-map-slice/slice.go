package main

import "fmt"

func main() {
	items := []int{10, 20, 30, 40, 50}
	for key, value := range items {
		fmt.Println(key, value)
	}
}
