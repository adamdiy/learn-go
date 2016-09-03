package main

import (
	"fmt"
	"strings"
)

//This is a copy of text.
func titleCase(text1, text2 string) (s1, s2 string) {
	text1 = strings.Title(text1)
	text2 = strings.Title(text2)
	return text1, text2
}

func main() {
	text1 := "Hello Don"
	text2 := "Hello You"
	fmt.Println(titleCase(text1, text2))
}
