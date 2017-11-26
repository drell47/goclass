package shouter

import (
	"regexp"
	"strings"
	"testing"
)

func TestShout(t *testing.T) {
	testString := "The quick brown fox"
	noLower := regexp.MustCompile("[a-z]")

	shoutedString := Shout(testString)
	if noLower.MatchString(shoutedString) {
		t.Fatalf("Found Lowercase")
	}
}

func TestShoutReallyLoud(t *testing.T) {
	testString := "The quick brown fox"
	noLower := regexp.MustCompile("[a-z]")

	shoutedString := ShoutReallyLoud(testString)
	if noLower.MatchString(shoutedString) {
		t.Fatalf("Found Lowercase")
	}
	final := len(shoutedString) - 1
	if shoutedString[final] != '!' {
		t.Fatalf("NO ! at end of shout")
	}
}

func TestReadAndShout(t *testing.T) {
	testString := "Hello\nWorld\n"
	s, err := ReadAndShout(strings.NewReader(testString))
	if err != nil {
		t.Fatalf("error reading test string ioReader")
	}
	noLower := regexp.MustCompile("[a-z]")
	if noLower.MatchString(s) {
		t.Fatalf("Found Lowercase")
	}
}
