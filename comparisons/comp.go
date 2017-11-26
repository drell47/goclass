package main

import (
	"fmt"
)

type MyString struct {
	Str string
}

func main() {

	setCase()

	s := MyString{Str: "Hello"}
	t := MyString{Str: "World"}

	fmt.Printf("%t\n", s == t)
	fmt.Printf("comper func - %t\n", comper(s, t))
}

func comper(s, t interface{}) bool {
	return s == t
}
