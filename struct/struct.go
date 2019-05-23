package main

import "fmt"

type person struct {
	name string
	age  uint8
}

func mainnn() {
	var p person

	p.name = "John Doe"
	p.age = 21

	// p := person{"John", 25}
	// p := person{age: 25, name: "John"}
	// p := struct{name string; age int}{"Amy",18} // anonymouse struct

	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
}
