package poetry

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Line is  declare types used for poetry Example
type Line string

// Stanza contains a group of lines in a Poem
type Stanza []Line

// Poem contains Stanza that make up a Poem
type Poem []Stanza

// NewPoem returns an empty Poem
func NewPoem() Poem {
	return Poem{}
}

// Add functions to Poem to allow sorting the stanzas of the poem
//   Need functions - Len - number of elements in the collection
//                    Less - if element i should be before element j
//                    Swap - swap element i and element j
func (s Stanza) Len() int {
	return len(s)
}
func (s Stanza) Less(i, j int) bool {
	return len(s[i]) <= len(s[j])
}
func (s Stanza) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	return
}

// NumStanzas returns the number of Stanzas in a Poem
func (p Poem) NumStanzas() int {
	return len(p)
}

// NumLines returns the number of Line in a Stanza
func (s Stanza) NumLines() int {
	return len(s)
}

// NumLines returns number of lines in a Poem
func (p Poem) NumLines() (count int) {
	for _, s := range p {
		count += s.NumLines()
	}
	return
}

// NumWords returns count of words (sep by space) in the poem
func (p Poem) NumWords() (count int) {
	// use strings.Split function to count words in poemFile
	for _, s := range p {
		for _, l := range s {
			count += len(strings.Split(string(l), " "))
		}
	}
	return
}

// LoadPoem will load a poem from a file
// example of using io to read from a files
func LoadPoem(name string) (Poem, error) {
	f, err := os.Open(name)
	if err != nil {
		// if get error, return nil and the error
		return nil, err
	}
	// set up for file to close at ends
	defer f.Close()

	p := Poem{}
	s := Stanza{}

	// use bufio Scanner to read lines from the files
	//   Scanner can specify different break points (other than new line...)
	scan := bufio.NewScanner(f)
	// read lines from the poem files
	for scan.Scan() {
		l := scan.Text()
		if len(l) == 0 {
			//s = append(s, "\n")
			p = append(p, s)
			s = Stanza{}
		} else {
			s = append(s, Line(l))
		}
		// fmt.Printf(" %d, %s\n", len(l), l)
	}

	return p, nil
}

// // String function performs formatting of the Poem object for Printing
// func (p Poem) String() string {
// 	result := ""
// 	for _, s := range p {
// 		for _, l := range s {
// 			result += fmt.Sprintf("%s\n", l)
// 		}
// 		result += "\n"
// 	}
// 	return result
// }

// String function performs formatting on a Stanzas
func (s Stanza) String() string {
	result := ""
	for _, l := range s {
		result += fmt.Sprintf("%s\n", l)
	}
	result += "\n"
	return result
}

// String function performs formatting on a NewPoem
func (p Poem) String() string {
	result := ""
	for _, s := range p {
		result += fmt.Sprintf("%s", s)
	}
	return result
}

// Stats returns number of vowels and consonants in a Poem
func (p Poem) Stats() (numVowels, numConsonants, NumPuncts int) {
	for _, stanza := range p {
		for _, line := range stanza {
			for _, r := range line {
				switch r {
				case 'a', 'e', 'i', 'o', 'u':
					numVowels++ // named return values already initialized
				case '.', ',', '?', '/', ':', ';', ' ':
					NumPuncts++
				default:
					numConsonants++ // named return values already initialized
				}
			}
		}
	}
	return
}

// func main() {
// 	fmt.Printf("Start main \n")
//
// }
