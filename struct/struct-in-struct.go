package main

import "fmt"

type Human struct {
	name string
	age  uint8
}

type Student struct {
	Human
	major string
}

func mainn() {
	john := Student{Human{"John", 21}, "IT"}
	fmt.Println(john.name, john.age, john.major)
}
