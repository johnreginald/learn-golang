package main

import "fmt"

func mainn() {
	items := map[string]int{"john": 20, "mary": 50}
	for key, value := range items {
		fmt.Println(key, value)
	}
}
