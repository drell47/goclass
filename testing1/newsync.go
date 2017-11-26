package main

import (
	"fmt"
)

func newmain() {
	fmt.Println("IN THE NEW MAIN")

	// channel for sending/receiving an integer
	c := make(chan int)
	// channel used to indicate program done, to stop the go routine
	done := make(chan struct{})

	go func() {
		i := 0

		for {
			select {
			case c <- i:
				i++
			case <-done:
				c <- 999
				return
			}
		}
	}()

	fmt.Printf("%d\n", <-c)
	fmt.Printf("%d\n", <-c)

	done <- struct{}{}
	fmt.Printf("%d\n", <-c)
}
