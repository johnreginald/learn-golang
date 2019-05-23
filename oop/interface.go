package main

import "fmt"

type Talkable interface {
	talk() string
}

type Human struct {
	name string
	age  int
	lang string
}

type American struct {
	Human
}

type Myanmar struct {
	Human
}

func (human *Human) talk() string {
	return "Hello World in " + human.lang + " by " + human.name
}

func test() {
	mm := Myanmar{Human{"Htet", 21, "Myanmar"}}

	fmt.Println(mm.name)
	fmt.Println(mm.age)
	fmt.Println(mm.talk())

	usa := American{Human{"John", 25, "English"}}

	fmt.Println(usa.name)
	fmt.Println(usa.age)
	fmt.Println(usa.talk())
}
