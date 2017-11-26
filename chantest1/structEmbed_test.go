package main

import (
	"regexp"
	"testing"
)

func TestShout(t *testing.T) {
	testString := "This Is a TEST string With CAPS 66"

	noLower := regexp.MustCompile("[a-z]")

	testOut := NewShoutingString(testString)

	if noLower.MatchString(testOut.str) {
		t.Fatalf("Found lower case letters")
	}

	if testOut.str != "THIS IS A TEST STRING WITH CAPS 66" {
		t.Fatalf("Shout string did not match")
	}
}
