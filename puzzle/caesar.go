package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(encrypt("Hello World", 5))
}

func encrypt(text string, salt uint8) string {
	newText := []uint8{}

	for _, element := range text {
		result := uint8(element) + salt

		if unicode.IsLower(rune(element)) {
			if result > 122 {
				newText = append(newText, result-26)
			} else {
				newText = append(newText, result)
			}
		}

		if unicode.IsUpper(rune(element)) {
			if result > 90 {
				newText = append(newText, result-26)
			} else {
				newText = append(newText, result)
			}
		}

		if element == 32 {
			newText = append(newText, 32)
		}
	}

	return string(newText)
}
