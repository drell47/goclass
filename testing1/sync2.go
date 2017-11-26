package main

import (
	"fmt"
)

// Counter - Another sync example using channels
type Counter struct {
	c    chan int
	done chan struct{}
	i    int
}

// NewCounter - create a new counter object and start the counter go function
func NewCounter() *Counter {
	counter := new(Counter)
	counter.c = make(chan int)
	counter.done = make(chan struct{})
	counter.i = 0

	// use go function to start counter - writes count to channel
	go func() {
		for {
			select {
			case counter.c <- counter.i:
				counter.i++
			case <-counter.done:
				counter.c <- 999
				return
			}
		}
	}()

	return counter
}

// GetSource - gets c chan from Counter object - returns read-only chan int
func (c *Counter) GetSource() <-chan int {
	return c.c
}

// Stop - send message to done channel in Counter - to stop counting go func
func (c *Counter) Stop() {
	c.done <- struct{}{}
}

func othermain() {
	fmt.Println("IN THE NEW MAIN")

	c := NewCounter()
	read := c.GetSource()

	fmt.Printf("%d\n", <-read)
	fmt.Printf("%d\n", <-read)

	// read is a read only channel  - so this gets a compile time Error
	// read <- 5

	c.Stop()
	fmt.Printf("%d\n", <-read)
}
