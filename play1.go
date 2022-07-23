package main

import (
	"fmt"
	"unicode"
)

func ToUpper(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToUpper(r[i])
	}
	return string(r)
}

func main() {
	fmt.Println(ToUpper("Hello"))
}
