package main

import (
	"fmt"
	"log"
	"net/http"
)

type Job struct {
	URL string
}
type Result struct {
	Status string
	URL    string
	Body   *http.Response
}

func crawl(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		resp, err := http.Get(job.URL)
		if err != nil {
			log.Fatal(err)
		}
		results <- Result{Status: resp.Status, URL: job.URL, Body: resp}
	}
}

func main() {
	jobs := make(chan Job, 3)
	results := make(chan Result, 3)

	// Spawn 5 worker goroutines
	for w := 1; w <= 7; w++ {
		go crawl(w, jobs, results)
	}

	// Send 5 jobs and close the channel
	urls := []string{
		"https://tutorialedge.net",
		"https://tutorialedge.net/pricing/",
		"https://example.com",
		"https://google.com",
		"https://amazon.com",
		"https://golang.org",
		"https://www.facebook.com",
	}

	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs)

	// Print results
	for r := 1; r <= 7; r++ {
		result := <-results
		fmt.Println(result)
	}
}
