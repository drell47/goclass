package main

// Interface example
//    interface defines capabilities, not types

import (
	"fmt"
	"math/rand"
)

type intshuf []int

func (s intshuf) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intshuf) Len() int {
	return len(s)
}

type stringslice []string

func (s stringslice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s stringslice) Len() int {
	return len(s)
}

// random shuffle of
type shuffler interface {
	// define capabilities required to satisfy shuffler Interface
	Len() int
	Swap(i, j int)
}

// define a function that operates on anything satisfying the shuffler Interface
func shuffle1(s shuffler) {
	// use Perm function from math/rand to define permutation
	num := int(s.Len())
	permutation := rand.Perm(num)
	// fmt.Printf("Perm= %v\n", permutation)
	for i := 0; i < num; i++ {
		if permutation[i] > i {
			s.Swap(i, permutation[i])
		}
	}
}

// define a function that operates on anything satisfying the shuffler Interface
func shuffle(s shuffler) {
	// use Perm function from math/rand to define permutation
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len() - i)
		// fmt.Printf("%d ", j)
		s.Swap(i, j)
	}
	// fmt.Printf("\n")
}

// define a function that operates on anything satisfying the shuffler Interface
func shuffle2(s shuffler) {
	// use Perm function from math/rand to define permutation
	for i := 0; i < s.Len()-1; i++ {
		j := rand.Intn(s.Len()-(i+1)) + (i + 1)
		// fmt.Printf("%d ", j)
		s.Swap(i, j)
	}
	// fmt.Printf("\n")
}

func main() {
	fmt.Printf("Start main \n")

	// keep count of which number ends in column 1
	col1 := []int{0, 0, 0, 0, 0}
	iter := 10000
	worig := intshuf{1, 2, 3, 4, 5}
	w := make(intshuf, len(worig))
	ind := 0
	for i := 0; i < iter; i++ {
		copy(w, worig)
		// fmt.Println(w)
		//shuffle(w)
		shuffle1(w)
		//shuffle2(w)
		///fmt.Println("<<<< ", w)
		ind = int(w[1]) - 1
		col1[ind]++
	}
	fmt.Printf("ind - %d\n", ind)
	fmt.Printf("%d tries - %v", iter, col1)

	ss := stringslice{"abc", "cde", "efg", "xyz", "qhr"}
	shuffle(ss)
	fmt.Printf("%v\n", ss)

}
