// Basic web server that returns a poem file
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/drell47/goclass/poetry"
)

var sError error
var poemFile string

var c = &config{}

// cache for pre-loading the valid ValidPoems
var cache map[string]poetry.Poem

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
		}

		poemFile := "/home/david/golang/" + poemName[0]
		newp, fileErr := poetry.LoadPoem(poemFile)
		if fileErr != nil {
			http.Error(w, "File not Found", http.StatusInternalServerError)
		} else {
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

	// read in configuration data from json file
	f, err := os.Open("/home/david/golang/config")
	if err != nil {
		fmt.Printf("Can not open config file")
		os.Exit(1)
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
