package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func main() {
	p := person{"John", 25}
	printInfo(p)
}

func printInfo(p person) {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
}
