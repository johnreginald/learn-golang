package main

import "fmt"

type math struct {
	x int
	y int
}

func (m math) add() int {
	return m.x + m.y
}

func (m math) subtract() int {
	return m.x - m.y
}

func (m math) multiply() int {
	return m.x * m.y
}

func (m math) division() int {
	return m.x / m.y
}

func main() {
	m := math{10, 20}
	fmt.Println(m.add())
	fmt.Println(m.subtract())
	fmt.Println(m.multiply())
	fmt.Println(m.division())
}
