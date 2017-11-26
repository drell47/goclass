package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/drell47/goclass/shouter"
)

func notmain() {
	testString := "abcdefg"
	fmt.Printf("Last char is g:  %c\n", testString[3])

	filename := flag.String("file", "", "File to Shout")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatalf("Failed to open file %s\n", *filename)
		return
	}

	s, err := shouter.ReadAndShout(file)
	if err != nil {
		fmt.Println("Error calling ReadAndShout")
	} else {
		fmt.Println(s)
	}
}
