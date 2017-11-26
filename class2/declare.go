package main

import (
	"fmt"
)

func empty(i []int) (n int) {
	// if len(i) > 0 {
	// Declaring this n in the if scope will shadow the return value 'n'
	if n := len(i); n > 0 {
		n = 1
	}
	return
}

func dectest() {

	var intSlice = []int{1, 2, 3, 4, 5, 6}
	stringSlice := []string{"this", "is", "a", "slice", "of", "string"}

	fmt.Printf("intSlice Empty Check: %d\n", empty(intSlice))

	if n := len(intSlice); n > -0 {
		fmt.Printf("intSlice has %d elements\n", n)
	} else {
		fmt.Printf("intSlice empty - %d\n", n)
	}

	if n := len(stringSlice); n > -0 {
		fmt.Printf("stringSlice has %d elements\n", n)
	} else {
		fmt.Printf("stringSlice empty - %d\n", n)
	}

	// test for error on redeclare variable
	n0, err := fmt.Printf("intSlice has %d elements\n", len(intSlice))
	n1, err := fmt.Printf("stringSlice has %d elements\n", len(stringSlice))
	fmt.Printf("n0,n1,err: %d. %d, %s\n", n0, n1, err)
}
