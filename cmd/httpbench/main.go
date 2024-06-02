// Define command line flags for benchmarking using the flag package
// Parse the command line arguments to ge the user-provided values for URL, number of requests, concurrency, and timeout
// Validate input and handle errors
package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	url         string
	numRequests int
	concurrency int
	timeout     time.Duration
)

// Define command line flags
func init() {
	flag.StringVar(&url, "url", "", "URL to benchmark")
	flag.IntVar(&numRequests, "n", 100, "Number of requests to send")
	flag.IntVar(&concurrency, "c", 1, "Number of requests to run concurrently")
	flag.DurationVar(&timeout, "timeout", 5*time.Second, "Request timeout")
}

func main() {
	flag.Parse()

	// Validate input
	if url == "" {
		fmt.Println("Error: URL is required")
		flag.Usage()
		os.Exit(1)
	}

	// Print parsed values
	fmt.Println("URL:", url)
	fmt.Println("Requests:", numRequests)
	fmt.Println("Concurrency:", concurrency)
	fmt.Println("Timeout:", timeout)

}
