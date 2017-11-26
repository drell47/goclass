package main

// use closing a channel to help coordinate functions

import (
	"fmt"
	"math/rand"
	"time"
)

// print a message after receiving on a channel (or channel closed)
func printer(msg string, goCh chan bool) {

	// wait to receive on the channel (don't care about value)
	<-goCh

	// print message after receiving on the channel
	fmt.Printf("%s\n", msg)

}

// keep printing until receive on channel, or channel is closed
func keepPrinting(msg string, stopCh chan bool) {

	// keep looping
	counter := 0
	for {
		// wait one second between prints
		time.Sleep(1 * time.Second)

		select {
		case <-stopCh:
			return
		default:
			fmt.Printf("%d: %s\n", counter, msg)
			counter++
		}
	}
}

// example - set channel to nil to prevent receiving messages on that channel
func reader(intCh chan int) {
	// start timer that will send on its channel in 3 seconds
	t := time.NewTimer(3 * time.Second)

	// loop forever
	for {
		// get the integer message or the timer message
		select {
		// message on integer channel
		case i := <-intCh:
			fmt.Printf("%d\n", i)
			// message from timer channel
		case <-t.C:
			intCh = nil // set int channel to nil - won't receive any more
		}
	}
}

func writer(ch chan int) {
	// write random integers to the int channel
	for {
		ch <- rand.Intn(42)
	}
}

func main() {

	fmt.Println("Start printers that wait for a GO")
	// create the channel used by the printer func
	goCh := make(chan bool)
	// create and start 10 copies of the printer function
	// use message that includes number of the copy
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("go channel printer %d", i), goCh)
	}
	// now wait 5 seconds
	time.Sleep(5 * time.Second)

	// if send on the channel, only one go routine would get it

	fmt.Println("Close the channel")
	// close the channel - all the go routines will stop waiting
	close(goCh)
	// now wait 5 seconds to let everything finish up
	time.Sleep(5 * time.Second)

	fmt.Println("\nStart printers that wait for a STOP")

	stopCh := make(chan bool)

	// create and start 10 copies of the printer function
	// use message that includes number of the copy
	for i := 0; i < 10; i++ {
		go keepPrinting(fmt.Sprintf("stop channel printer %d", i), stopCh)
	}

	// now wait 5 seconds
	time.Sleep(5 * time.Second)

	// if send on the channel, only one go routine would get it

	fmt.Println("Close the channel")
	// close the channel - all the go routines will stop waiting
	close(stopCh)

	// now wait 5 seconds to let everything finish up
	time.Sleep(5 * time.Second)

	// Test go routines that set channel to nil
	intCh := make(chan int)
	fmt.Println("\nStarting reader and writer for nil channel test\n")
	go reader(intCh)
	go writer(intCh)

	time.Sleep(10 * time.Second)

}
