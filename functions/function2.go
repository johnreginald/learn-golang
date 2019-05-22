package main

import "fmt"

func main() {
	test(10, 20, 30, 40, 50)
}

func test(arg ...int) {
	fmt.Println(arg)
}
