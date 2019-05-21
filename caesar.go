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

	for x := 0; x < len(text); x++ {
		result := text[x] + salt

		if unicode.IsLower(rune(text[x])) {
			if result > 122 {
				newText = append(newText, result-26)
			} else {
				newText = append(newText, result)
			}
		}

		if unicode.IsUpper(rune(text[x])) {
			if result > 90 {
				newText = append(newText, result-26)
			} else {
				newText = append(newText, result)
			}
		}

		if text[x] == 32 {
			newText = append(newText, 32)
		}
	}

	return string(newText)
}
