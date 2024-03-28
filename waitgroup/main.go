package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://www.google.com",
	"https://www.facebook.com",
	"https://www.twitter.com",
}

func fetch(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	// Decrement the counter when the goroutine completes.
	wg.Done()
}

func crawl() {
	// Create a new WaitGroup.
	var wg sync.WaitGroup

	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)

		go fetch(url, &wg)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
	fmt.Println("All fetches completed")
}

func main() {
	crawl()
}
