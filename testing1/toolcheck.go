package main

import (
	"flag"
	"fmt"
	"time"
)

// Printf - dummy local function
func Printf(s string) {
	fmt.Println(s)
}

func main() {
	timeout := flag.Duration("timeout", 1*time.Millisecond, "timeout value")
	flag.Parse()

	fmt.Printf("Timeout value is %s\n", *timeout)

	fmt.Printf("Timeout int is %d\n", *timeout)
	//t1, err := time.ParseDuration(timeout)
	//if err != nil {
	//	fmt.Printf("error parsing timeout: %s\n", err)
	//} else {
	//	fmt.Printf("t1 is %v\n", t1)
	//}

}
