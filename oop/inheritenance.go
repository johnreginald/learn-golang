package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // anonymous field
	school string
}

type Employee struct {
	Human
	company string
}

// define a method in Human
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}

	sam.SayHi()
	mark.SayHi()
}
