// Basic web server that returns a poem file
package main

import (
	"encoding/json"
	"flag" // for command line arguments
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/drell47/goclass/poetry"
)

var sError error
var poemFile string

var c = &config{}

// cache for pre-loading the valid ValidPoems
// var cache = make(map[string]poetry.Poem)

// create a Mutex to protect the cache
// var cacheMutex sync.Mutex

// Alternate way to apply mutex to a struct
type protectedCache struct {
	sync.Mutex
	c map[string]poetry.Poem
}

var cache protectedCache

// struct to use for configuration data for program
//    read in using json data  - Route is url
//        BindAddress is port to use - give json hint that it's called addr
type config struct {
	Route       string
	BindAddress string   `json:"addr"`
	ValidPoems  []string `json:"valid"`
}

// structure to use with testinging encoding/json
//   only public (capital first letter) will be in json encoding
type poemWithTitle struct {
	Title    string
	Body     poetry.Poem
	NumWords string // use strconv to convert int to string
}

func readPoem(poemName string) (poetry.Poem, error) {

	poemFile := "/home/david/golang/" + poemName
	newp, fileErr := poetry.LoadPoem(poemFile)
	if fileErr != nil {
		return nil, fileErr
	}
	return newp, nil
}

// poemhandler responsds to a request and writes back selected poem
func poemHandler(w http.ResponseWriter, r *http.Request) {

	// get information from the request
	//   ParseForm - populates r.Form and r.PostForm
	r.ParseForm()

	// Check if use specify poem name in URL  (?name=xxxxx)
	//    returns a slice of string
	poemName, ok := r.Form["name"]
	if !ok {
		fmt.Fprintf(w, "No name supplied, try ?name=  poe or chaucer")
	} else {
		// check valid poem names
		found := false
		for _, name := range c.ValidPoems {
			if poemName[0] == name {
				found = true
			}
		}
		if !found {
			http.Error(w, "INVALID POEM NAME", http.StatusNotFound)
		} else {

			log.Printf("User selected poem %s\n", poemName[0])

			//   Switch to using poems from the cache - so only load once
			newp, ok := cache.c[poemName[0]]
			if !ok {
				http.Error(w, "Not Found (invalid)", http.StatusNotFound)
				return
			}
			// poemFile := "/home/david/golang/" + poemName[0]
			// newp, fileErr := poetry.LoadPoem(poemFile)
			// if fileErr != nil {
			// 	http.Error(w, "File not Found", http.StatusInternalServerError)
			// } else {
			// write out using String formatting to the ioWriter
			// fmt.Fprintf(w, "%s\n", newp)

			// sort the poem (shorted line to front of stanza) - do first stanza
			sort.Sort(newp[0])

			numWords := int64(newp.NumWords())
			strNumWords := strconv.FormatInt(numWords, 16)
			fmt.Fprintf(w, "Poem: %s, has 0x%s words\n", poemName[0], strNumWords)
			pwt := poemWithTitle{poemName[0], newp, "0x" + strNumWords}

			enc := json.NewEncoder(w)
			enc.Encode(pwt)

		}
	}

}

func main() {

	// set up config file with
	//  {
	//    "Route" : "http://127.0.0.1",
	//    "addr" : "8080"
	//  }

	// Configure logging  (can also log to file instead of stdout)
	// log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)
	log.SetFlags(log.Ldate | log.Ltime)

	// see log/syslog for remote manage logs

	// use flag package to get command line arg for config file
	//    flag.String and flag.StringVar are for cmd line flags    -xyz=abc
	//       NOTE: can use flag.Arg()  to get command line arguments also
	defaultConfig := "/home/david/golang/config"

	// define flags on command line  - flag.String returns a pointer to string
	//configFilePtr := flag.String("conf", defaultConfig, "Path to configuration file")

	// alternate way - flag.StringVar updates a string using a supplied pointer
	var configFileStr string
	flag.StringVar(&configFileStr, "conf", defaultConfig, "Path to configuration file")

	// parse data from defined flags from command line
	flag.Parse()

	// read in configuration data from json file

	// from flag.String, get a pointer to string
	// f, err := os.Open(*configFilePtr)

	// from flag.StringVar, use the updated String variable
	f, err := os.Open(configFileStr)

	if err != nil {
		// fmt.Printf("Can not open config file")
		log.Fatalf("Failed to open config file \n")
		//os.Exit(1)
	}

	dec := json.NewDecoder(f) // f is an ioReader
	// cptr := &c
	err = dec.Decode(c)
	if err != nil {
		fmt.Printf("decode error : %s\n", err)
	}
	f.Close()

	fmt.Printf("Route: %s, BindAddress: %s\n", c.Route, c.BindAddress)
	fmt.Printf("Valid poem names: %v\n", c.ValidPoems)

	// // Load the poems into the cache
	// for _, vp := range c.ValidPoems {
	// 	newp, perr := readPoem(vp)
	// 	if perr != nil {
	// 		fmt.Printf("Error loading cache for %s\n", vp)
	// 		os.Exit(1)
	// 	}
	// 	cache[vp] = newp
	// }

	// set up WaitGroup to ensure wait for all poem reads to finish
	var wg sync.WaitGroup
	cache.c = make(map[string]poetry.Poem)

	// Check how much time it takes to load the ValidPoems
	startTime := time.Now()

	// Load the poems into the cache using go routines
	for _, vp := range c.ValidPoems {
		// add 1 to wait group count
		wg.Add(1)
		go func(n string) {
			// use the anonymous mutex in the cache
			cache.Lock() // cacheMutex.Lock()
			// defer cacheMutex.Unlock()     // if have many lines to keep locked
			cache.c[n], err = readPoem(n) // go routines sharing a map - need sync
			cache.Unlock()                // cacheMutex.Unlock()
			if err != nil {
				fmt.Printf("Error loading cache for %s\n", n)
				log.Fatalf("Error loading cache for %s\n", n)
				// os.Exit(1)
			}
			wg.Done() // could use defer
		}(vp)
	}

	wg.Wait() // wait until all go routines finish reading poems

	// Should wait for all the go routines reading poems to finish before
	//  going ahead with the http server
	//   Use a waitgroup from the sync packages

	var elapsed time.Duration
	//elapsed = time.Since(startTime)

	// alternate
	stopTime := time.Now()
	elapsed = stopTime.Sub(startTime)

	log.Printf("Loading took: %s\n", elapsed)
	// log.Printf("Started at %s\n", startTime.Format(time.Kitchen))
	log.Printf("Started at %s\n", startTime.Format(time.RFC822))
	// specify handler for requested url - adds handler to DefaultServeMux
	//   func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	base := fmt.Sprintf("%s", c.Route)
	port := fmt.Sprintf("%s", c.BindAddress)
	http.HandleFunc(base, poemHandler)

	// start listener that will serve web requests
	//    func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	//    handler - nil, means use DefaultServeMux
	hsError := http.ListenAndServe(port, nil)
	sError = hsError
}
