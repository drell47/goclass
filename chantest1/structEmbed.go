package main

import (
	"fmt"
	"strings"
	"sync"
)

// Example with EMBEDDING a struct in another struct
//    struct 'inherits' fields and methods from the embedded structs

// MOST COMMON USAGE:   - embed a mutex in a struct.
//   Can then use Lock() and Unlock() methods

// MyString is a struct used for embedded struct example
type MyString struct {
	sync.Mutex // embed Mutex
	str        string
}

// NewMyString function creates a new MyString struct
func NewMyString(s string) MyString {
	return MyString{str: s}
}

// Output function will print out the MyString struct
func (m MyString) Output() {
	fmt.Println(m.str)
}

// Shouting is struct with embedded struct
type Shouting struct {
	MyString        // embedded struct
	junk     string // new field
}

// Output is a method on Shouting struct
func (s Shouting) Output() {
	fmt.Printf("Shouting Output: %s\n", s.junk)
}

// NewShoutingString - makes struct with string as upper case
func NewShoutingString(s string) Shouting {
	load := Shouting{}
	load.str = strings.ToUpper(s)
	return load
}

func structEmbedRunit() {
	hello := NewMyString("Embed Check - Hello There")
	shout := NewShoutingString("Shout Embed Check - Hello There")
	hello.Output()
	shout.Output()
	shout.junk = "This is Junk - OK junk"
	shout.Output()
	fmt.Printf("shout: %v\n", shout)
}
