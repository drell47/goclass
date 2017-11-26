package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

var (
	// counter for how many things running at once
	running int64 = 0
)

// print with waiting
func work() {
	// increment count of running work functions
	atomic.AddInt64(&running, 1)
	fmt.Printf("[%d", running)
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	// decrement count of running functions
	atomic.AddInt64(&running, -1)
	fmt.Printf("]")
}

func worker(sema chan bool) {
	// wait for semaphore
	<-sema
	// perform work
	work()
	// signal semaphore that another worker can go
	sema <- true
}

func main() {
	// set up semaphore channel - buffered to hold up to 10 messages
	sema := make(chan bool, 10)
	fmt.Println("Start main")
	// start 1000 worker go functions
	for i := 0; i < 1000; i++ {
		go worker(sema)
	}

	// send out 10 messages on the semaphore channel
	for i := 0; i < cap(sema); i++ {
		sema <- true
	}

	time.Sleep(10 * time.Second)

	// Unbuffered channel can only read or write if there is something ready
	// ch := make(chan int)

	// buffered channel can buffer up to specified size limit
	//   once it is full writes will wait
	// bufch := make(chan int, 100)

}
