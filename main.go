package main

import "fmt"

func main() {
	items := map[string]int{"john": 20, "mary": 50}
	for key, value := range items {
		fmt.Println(key, value)
	}
}

// func test(arg ...int) {
// for _; item := range arg {
// 	fmt.Println(item)
// }
// }
