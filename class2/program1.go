package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Print("Hello There\n")

	// call function in other file
	dectest()

	// example for break and continue and named loops
MyLabel:
	for {
		break MyLabel
	}

	var (
		even  int
		odd   int
		total int
		zeros int
	)

	numbers := []int{1, 2, 3, 5, 0, 7, 8, 9, -3, 10}

OuterLabel:
	for i := 0; i < 10; i++ {
		for _, n := range numbers {
			total++
			// skip over any 0 values
			// if n == 0 {
			//  	continue  // OuterLabel    // skip 0 value
			// }

			// stop once hit a 0
			if n == 0 {
				break OuterLabel // stop counting if hit a 0 - exit outer loop label
			}

			if n%2 == 0 {
				even++
			} else {
				odd++
			}
		}
	}
	fmt.Printf("Total: %d, Even: %d, Odd: %d\n", total, even, odd)

	var err error

	// Example using Switch
	even, odd, zeros, total = 0, 0, 0, 0
LoopTop:
	for _, n := range numbers {
		// switch statement doesn't fall through unless have "fallthrough" statement
		switch {
		case n == 0:
			zeros++
		case n%2 == 0:
			even++
		case n%2 == 1:
			odd++
		default:
			err = errors.New("Found Negative number")
			break LoopTop // negative mod 2 is still negative (not 0 or 1)
			// could also use a "continue",
		}
		total++
	}
	if err != nil {
		fmt.Printf("Found Negative Number\n")
	}
	fmt.Printf("Total: %d, Even: %d, Odd: %d\n", total, even, odd)
}
