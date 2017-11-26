package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/drell47/goclass/yelling"
)

type T1 struct {
	name string
}

type T2 struct {
	name string
}

func main() {
	fmt.Printf("This is main\n")

	var d1 = T1{name: "MyName"}
	var d2 T2
	d2 = T2(d1)
	fmt.Printf("d2 = %v\n", d2.name)

	ls := yelling.NewLoudString()
	ls.Change("Hello")
	fmt.Println(ls.String())

	n := yelling.NewLoudString // assign the method
	ls2 := n()

	change := ls2.Change
	change("Be quiet")
	fmt.Println(ls2.String())

	var _ yelling.Yeller = ls // added Len method, so would implement Yeller
	var _ fmt.Stringer = ls
}

// function to make some stuff for garbage collection
func length(s string) int {
	c := []byte(s)
	n := len(c)
	b := make([]byte, n)
	return len(b)
}

func garbcoll() {
	s := "Hello, World!"
	s += s
	s += s
	s += s
	fmt.Printf("Length of s: %d\n", length(s))

	start := time.Now()
	for {
		if time.Since(start) > time.Second {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Printf("Heap Size: %d\n", r.HeapAlloc)
			fmt.Printf("NumGC %d\n", r.NumGC)
			start = time.Now()
		}
		length(s)
	}

}

func used_to_be_the_main() {

	go garbcoll()
	time.Sleep(10 * time.Second)

	nums := []int{1, 11, 21, 1211, 111221, 312211}
	middle := nums[1:3]
	fmt.Printf("Middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))

	nums[1] *= 2
	fmt.Printf("Middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))

	middle = append(middle, 42)
	fmt.Printf("Middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))
	fmt.Printf("nums: %v, %d, %d\n", nums, len(nums), cap(nums))

	// add elements to middle to exceed capacity - creates new underlying array
	middle = append(middle, 43)
	middle = append(middle, 44)
	middle = append(middle, 45)
	middle = append(middle, 46)
	fmt.Printf("Middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))
	fmt.Printf("nums: %v, %d, %d\n", nums, len(nums), cap(nums))

	nums[1] *= 2
	fmt.Printf("Middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))
	fmt.Printf("nums: %v, %d, %d\n", nums, len(nums), cap(nums))

	// try to remove an element (say 43) - append 2 pieces without that one
	//    NOTE:   use three dots ('...') to pass individual elemtents to append
	middle = append(middle[:3], middle[4:]...)
	fmt.Printf("Middle: %v, len: %d, cap: %d\n", middle, len(middle), cap(middle))

}
