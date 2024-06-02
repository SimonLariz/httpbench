// Define command line flags for benchmarking using the flag package
// Parse the command line arguments to ge the user-provided values for URL, number of requests, concurrency, and timeout
// Validate input and handle errors
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/SimonLariz/httpbench/pkg/benchmark"
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

	// Call benchmarking function
	m := benchmark.Run(url, numRequests, concurrency, timeout)
	// benchmark.Run(url, numRequests, concurrency, timeout)
	fmt.Printf("Total requests: %d\n", m.TotalRequests)
	fmt.Printf("Successful requests: %d\n", m.SuccessfulRequests)
	fmt.Printf("Failed requests: %d\n", m.FailedRequests)
	fmt.Printf("Average response time: %v\n", m.AverageResponseTime())
	fmt.Printf("Throughput: %.2f req/s\n", m.RequestsPerSecond())
}
