package main

import (
	"fmt"
)

const (
	message3 = "The answer to all is %d\n"
	answer = 42
)


func main() {
	var message string
	message = "Hello, Jupiter\n"
	message2 := "Hello, Saturn\n"
	fmt.Println("Hello, World...")
	fmt.Print(message)
	fmt.Println(message2)
	//  answer += 1   // can't assign to constant
	fmt.Printf(message3, answer)
	
	pi := float64(3.14)
	fmt.Printf("%8.2f\n", pi)
	blocks := 421
	fmt.Printf("Blocks = %12d, and %-8d\n", blocks, blocks)
	
	istrue := true
	fmt.Printf("is true = %t\n", istrue)
}
