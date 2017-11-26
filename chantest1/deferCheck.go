package main

import (
	"errors"
	"fmt"
)

func do() {

	var err error
	defer func() {
		if err != nil {
			fmt.Printf("Error occurred: %s\n", err)
		}
	}()

	n := 1
	defer fmt.Printf("First: %d\n", n)
	n++
	defer fmt.Printf("Second: %d\n", n)
	n++
	defer fmt.Printf("Third: %d\n", n)

	fmt.Println("do function - do somthing")

	err = errors.New("Some trouble at the end")

	// could use fmt.Errorf()
}
