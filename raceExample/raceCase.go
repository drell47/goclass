package main

// RACE condition example -  run with   go run -race raceCase.go
//
//    to see race condition identified

// Also check for flag -version to print a version number

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var Version string = "1.301"

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

var start time.Time
var t *time.Timer

// functin that resets the timer to a random duration
func reseter() {
	fmt.Println(time.Now().Sub(start))
	t.Reset(randomDuration())
}

func main() {

	version := flag.Bool("version", false, "Get version")
	flag.Parse()
	if *version {
		fmt.Printf("Program Version: %s\n", Version)
		return
	}

	fmt.Println("HI THERE")
	start = time.Now()
	//var t *time.Timer

	t = time.AfterFunc(randomDuration(), reseter)

	// wait to give function time to run
	time.Sleep(5 * time.Second)
}
