package writers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// func Get(url string) (resp *Response, err error)

// get the size of web page at url, int and possible error
func getPage_orig(url string) (int, error) {
	fmt.Printf("getPage called\n")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return 0, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error 2: %s", err)
		return 0, err
	}

	return len(body), nil
}

func test_main() {
	urls := []string{"http://www.google.com", "http://www.yahoo.com", "http://www.bing.com", "http://www.bbc.co.uk"}
	for _, url := range urls {
		plen, err := getPage_orig(url)
		if err != nil {
			fmt.Printf("Error calling getPage:, %s", err)
			os.Exit(1)
		} else {
			fmt.Printf("%s is Length = %d\n", url, plen)
		}
	}
}
