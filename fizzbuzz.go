package main

import "fmt"

func main() {
	for x := 1; x <= 100; x++ {
		if x%3 == 0 && x%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if x%3 == 0 {
			fmt.Println("Fizz")
		} else if x%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(x)
		}
	}
}
