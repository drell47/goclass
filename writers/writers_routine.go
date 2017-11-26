package writers // main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
)

// func Get(url string) (resp *Response, err error)

// get the size of web page at url, int and possible error
func getPage(url string) (int, error) {
	fmt.Printf("getPage called: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %d", err)
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error 2: %d", err)
		return 0, err
	}

	return len(body), nil
}

// function to call as go routine to do calls to getPage
//     use channel to return size value
func getter(url string, size chan int) {
	length, err := getPage(url)
	if err == nil {
		size <- length
	}
	close(size)

}

func main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com", "http://www.bbc.co.uk"}
	num_urls := len(urls)
	chans := make([]chan int, num_urls)
	for i, url := range urls {
		chans[i] = make(chan int)
		go getter(url, chans[i])
	}
	not_done := true
	answers := 0
	plen := 0
	for not_done {
		select {
		case plen = <-chans[0]:
			if plen > 0 {
				fmt.Printf("%s is Length = %d\n", urls[0], plen)
				answers++
			}
		case plen = <-chans[1]:
			if plen > 0 {
				fmt.Printf("%s is Length = %d\n", urls[1], plen)
				answers++
			}
		case plen = <-chans[2]:
			if plen > 0 {
				fmt.Printf("%s is Length = %d\n", urls[2], plen)
				answers++
			}
		case plen = <-chans[3]:
			if plen > 0 {
				fmt.Printf("%s is Length = %d\n", urls[3], plen)
				answers++
			}
		default:
			// fmt.Println("nothing")
		}
		if answers >= 4 {
			not_done = false
		}
	}

}
