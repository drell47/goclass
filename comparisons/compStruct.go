package main

import (
	"fmt"
)

type Person struct {
	First, Last string
}

// set up a map with anonymous struct to represent a set (present or not)

// can also use a bool value and use ok to see if key present...

func setCase() {

	people := make(map[Person]struct{})

	js := Person{"John", "Smith"}
	jb := Person{"Justin", "Bieber"}
	as := Person{"Adam", "Smith"}

	people[js] = struct{}{}
	people[jb] = struct{}{}

	fmt.Printf("%#v\n", people)

	_, ok := people[as]
	fmt.Printf("%s %s is in people = %t\n", as.First, as.Last, ok)

}
