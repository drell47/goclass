package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var numWorkers int

type task interface {
	process()
	output()
}

type factory interface {
	create(line string) task
}

func run(f factory) {
	var wg sync.WaitGroup
	begin = time.Now()

	in := make(chan task)

	wg.Add(1)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			in <- f.create(s.Text())
		}
		if s.Err() != nil {
			log.Fatalf("Error reading STDIN: %s", s.Err())
		}
		close(in)
		wg.Done()
	}()

	out := make(chan task)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			for t := range in {
				t.process()
				out <- t
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for t := range out {
		t.output()
	}
}

type HTTPTask struct {
	url     string
	ok      bool
	start   time.Duration
	elapsed time.Duration
}

func (h *HTTPTask) process() {

	if h.url == "" {
		h.ok = false
		return
	}
	start := time.Now()
	h.start = time.Since(begin)
	resp, err := http.Get(h.url)
	h.elapsed = time.Since(start)
	if err != nil {
		h.ok = false
		return
	}
	if resp.StatusCode == http.StatusOK {
		h.ok = true
		return
	}
	h.ok = false

}

func (h *HTTPTask) output() {
	fmt.Printf("%s %s %t %s\n", h.start, h.url, h.ok, h.elapsed)
}

type Factory struct {
}

func (f *Factory) create(line string) task {
	h := &HTTPTask{}
	h.url = line
	return h
}

var begin time.Time

func main() {
	fmt.Printf("Hello, World!\n")
	n := flag.Int("n", 10, "number of workers")
	flag.Parse()
	numWorkers = *n

	f := &Factory{}

	run(f)
}
