package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "123456.2545451"

	s = commaFunc(s)

	fmt.Println(s)
}

func comma(s string) string {

	if len(s) <= 3 {
		return s
	}

	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func commaFunc(s string) string {
	i := strings.LastIndex(s, ".")
	str1 := s[:i]
	str2 := s[i+1:]
	fmt.Printf("str1 = %q str2 = %q\n", str1, str2)
	for i := len(str1) - 3; i > 0; i -= 3 {
		str1 = str1[:i] + "," + str1[i:]

	}

	for i := len(str2) - 3; i > 0; i -= 3 {
		str2 = str2[:i] + "," + str2[i:]

	}

	return str1 + "." + str2
}
