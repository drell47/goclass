package main

import (
	"fmt"

	"github.com/drell47/goclass/poetry"
)

func main() {
	fmt.Printf("Start of Main")
	p := poetry.NewPoem()
	s := poetry.Stanza{}
	l := poetry.Line("hello there")

	LL := poetry.Stanza{"Who is there", "This is a test", "try to find words"}

	pt := poetry.Poem{{"Line 1 adg, frog?", "Line 2 up/down, uiop.", "Line 3 trewgfd,", "Line 4, bvcxde."}}
	pstanza := poetry.Poem{{"Line 1 adg, frog?", "Line 2 up/down, uiop."}, {"Line 3 trewgfd,", "Line 4, bvcxde."}}
	s = append(s, l)
	p = append(p, s)
	p = append(p, LL)

	fmt.Printf("p len = %d\n", len(p))
	v, c, pun := p.Stats()
	fmt.Printf("Stats are: vowels: %d, consonants: %d, punctuation: %d\n", v, c, pun)
	fmt.Printf("Stanzas: %d, Lines: %d\n", p.NumStanzas(), p.NumLines())
	fmt.Printf("PT: Stanzas: %d, Lines: %d\n", p.NumStanzas(), p.NumLines())

	v, c, pun = pt.Stats()
	fmt.Printf("PT: Vowels: %d, Consonants: %d, punctuation: %d\n", v, c, pun)
	fmt.Printf("\n")
	// Printing Example
	fmt.Printf("%s\n", pstanza)

	poemFile := "/home/david/golang/junk_poem.txt"
	newp, newerr := poetry.LoadPoem(poemFile)
	// fmt.Printf("new poem: %#v, New Err: %s", newp, newerr)
	fmt.Printf("new poem: %s, New Err: %s", newp, newerr)
}
