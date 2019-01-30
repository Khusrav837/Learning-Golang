package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "Зоб"
	str2 := "боз"
	b := anogramma(str1, str2)

	fmt.Printf("Natija = %t", b)
}

func anogramma(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	r1 := []rune(strings.ToUpper(str1))
	r2 := []rune(strings.ToUpper(str2))

	fmt.Printf("len = %d\n", len(str1))

	for i := 0; i < len(r1); i++ {

		fmt.Printf("str1 = %c str2 = %c \n", r1[i], r2[len(r2)-i-1])

		if r1[i] != r2[len(r2)-i-1] {
			return false
		}

	}

	return true
}
