package main

import "fmt"

type one func() string

func main() {
	one := func() string {
		return "Hello World"
	}
	two(one)
}

func two(f func() string) {
	fmt.Println(f())
}
