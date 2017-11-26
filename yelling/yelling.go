package yelling

import (
	"strings"
)

// Yeller is an interface that specifies methods implemented
type Yeller interface {
	String() string
	Change(s string)
	Blank()
	Len() int
}

// LoudString - type with string in Upper Case
type LoudString struct {
	// claim implements Yeller interface
	Yeller
	s string
}

// NewLoudString - return pointer to a LoudString
func NewLoudString() *LoudString {
	return &LoudString{}
}

// String - return the string value in a LoudString
func (ls *LoudString) String() string {
	return ls.s
}

// Change - change string value in LoudString to passed in string value
func (ls *LoudString) Change(s string) {
	ls.s = strings.ToUpper(s)
}

// Blank - replace string in LoudString with empty string
func (ls *LoudString) Blank() {
	ls.s = ""
}

// Len - return length of string element of LoudString
func (ls *LoudString) Len() int {
	return len(ls.s)
}
