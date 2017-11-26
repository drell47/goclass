package main

import (
	"fmt"
	"time"
)

// error type - interface with one method Error() string

// MyError is an error type (by virture of having Error() string method)
type MyError struct {
	str  string
	when time.Time
}

func (m MyError) Error() string {
	return fmt.Sprintf("%s at %s", m.str, m.when.Format(time.RFC822))
}

func errCheck() {
	err := MyError{str: "Bad Result", when: time.Now()}

	fmt.Printf("Error - %s\n", err)
}
