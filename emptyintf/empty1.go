package main

import (
	"fmt"
)

func whatIsThis1(i interface{}) {
	fmt.Printf("Type is %T\n", i)
}
func whatIsThis(i interface{}) {
	// type switch - execute case based on type of data
	//    passed in using the empty interface
	switch v := i.(type) {
	case string:
		fmt.Printf("It is a string: %s\n", v)
	case uint32:
		fmt.Printf("It is unsigned 32 bit integer: %d\n", v)
	default:
		fmt.Printf("Something else: %v\n", v)
	}
}

func typeAssert(i interface{}) {
	fmt.Printf("claim it is string: %s\n", i.(string))
}

func main() {
	fmt.Printf("Start main \n")

	i1 := 6
	f1 := 3.456
	s1 := "this is a test"
	a1 := [...]string{"aa", "bb", "cc"}
	whatIsThis(i1)
	whatIsThis(f1)
	whatIsThis(s1)
	whatIsThis(uint32(i1))
	whatIsThis(a1)

	// typeAssert(43)    // fails type assertion
	typeAssert("hello") // ok type assertion
}
