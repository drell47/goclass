package main

import (
	"fmt"
	"log"
	"regexp"
)

// use MustCompile to cause panic if regexp doesn't compile, don't check error
var matchNumber = regexp.MustCompile("[0-9]+")

func matchRunit() {
	re, err := regexp.Compile("[0-9]+")
	if err != nil {
		log.Fatalf("Bad regexp: %s", err)
	}
	phrase := "Hello"
	fmt.Printf("Does it match? (%s), %t\n", phrase, re.Match([]byte(phrase)))
	fmt.Printf("Does it match? (%s), %t\n", phrase, re.MatchString(phrase))

	// use the regexp from the MustCompile
	phrase2 := "Hello 31415"
	fmt.Printf("Does it match? (%s), %t\n", phrase2, matchNumber.MatchString(phrase2))
}
