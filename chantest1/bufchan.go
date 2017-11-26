package main

import (
	"fmt"
	"time"
)

// Buffer - example with a struct and two functions on the struct
//   one to get a struct, and one to do something (process it)
type Buffer struct {
	id int
}

func (b Buffer) process() {
	fmt.Printf("process id = %d\n", b.id)
	time.Sleep(200 * time.Millisecond)
}

func (b Buffer) get(source string) {
	fmt.Printf("get - id = %d,  source = %s\n", b.id, source)
	time.Sleep(400 * time.Millisecond)
}

// buffered channel that can hold up to 10 Buffer structs
var available = make(chan Buffer, 10)

// channel for passing one Buffer to the controller
var toController = make(chan Buffer)

// worker will get a Buffer from the available buffer chan or make a new one
func worker() {
	var source string
	for i := 0; i < 5; i++ {
		var b Buffer
		select {
		case b = <-available:
			source = "channel"
		default:
			b = Buffer{id: i}
			source = "new"
		}
		b.get(source)
		toController <- b
	}
}

// controller will do some processing on a buffer
func controller() {
	for {
		b := <-toController
		b.process()

		select {
		case available <- b:
		default:
			// just drop out of the select - Buffer will be discarded
		}
	}
}

func main() {

	// use of custom error tyep
	errCheck()

	// check on use of defer
	do()

	matchRunit()
	//go controller()
	//go worker()
	//time.Sleep(10 * time.Second)

	structEmbedRunit()

	interfaceRunit()
}
