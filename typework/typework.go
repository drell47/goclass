// Example with user defined types
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// user type defined using a struct
type webPage struct {
	url  string
	body []byte
	err  error
}

// user type defined from a slice of int
type SummableSlice []int

// define function on SummableSlice type
//    slice is passed as reference already, so don't need to use pointer
func (s SummableSlice) sum() int {
	sum := 0
	for _, val := range s {
		sum += val
	}
	return sum
}

// add a function tied to the webPage types
//   here, receiver is pointer to webPage type, so can modify the values
func (w *webPage) get() {
	resp, err := http.Get(w.url)
	if err != nil {
		w.err = err
	} else {
		w.body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			w.err = err
		}
	}
}

// add another function with reciever as a webPage item (not pointer), so
//    get a copy of webPage item, can't modify values in original item
func (w webPage) isOK() bool {
	return w.err == nil
}

func main() {

	// alternate ways to create webPage type item

	w := &webPage{url: "http://www.oreilly.com"}

	// w := webPage{}
	// w.url = "http://www.oreilly.com/"

	// w := webPage{url: "http://www.oreilly.com/"}

	// w := new(webPage)
	// w.url = "http://www.oreilly.com/"

	w.get()

	if w.isOK() {
		fmt.Printf("URL: %s, Error: %s, Length: %d\n", w.url, w.err, len(w.body))
	} else {
		fmt.Printf("Error with URL - %s, is %s", w.url, w.err)
	}

	// example using SummableSlice type
	s1 := SummableSlice{1, 2, 3, 4, 5}
	sum := s1.sum()
	fmt.Printf("Sum is %d\n", sum)

}
