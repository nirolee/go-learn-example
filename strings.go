package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "go好像不难"
	fmt.Println(utf8.RuneCountInString(s))

	for i, ch := range []rune(s) {
		fmt.Printf("%d %c", i, ch)
	}
}
