// Package shouter contains methods for uppercase strings
package shouter

import (
	"bufio"
	"io"
	"strings"
)

// A StringUpper is a type containing the lowercase and uppercase versions
type StringUpper struct {
	lower string
	upper string
}

// Shout - return uppercase string
func Shout(s string) string {
	return strings.ToUpper(s)
}

// ShoutReallyLoud - return uppercase with exclaimation
func ShoutReallyLoud(s string) string {
	return Shout(s) + "!"
}

// ReadAndShout - read lines from an io.Reader and Shout them
//  use an io.Reader instead of os.File - makes testing easier
//  and see if godoc puts this on different lines?
//    or if this is separated with blank line
func ReadAndShout(r io.Reader) (string, error) {
	gather := ""
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		gather += Shout(scanner.Text()) + "\n"
		// fmt.Println(Shout(scanner.Text()))
	}
	return gather, scanner.Err()
}
