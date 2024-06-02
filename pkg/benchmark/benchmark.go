// Core benchmarking functionality
// Function takes URL and sends HTTP requests to the URL using net/http package
// Implement logic for response time calculation
// Use goroutines to send multiple requests concurrently
package benchmark

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Run(url string, numRequests, concurrency int, timeout time.Duration) {
	// Implement benchmarking logic
	var wg sync.WaitGroup
	var totalDuration time.Duration

	client := &http.Client{
		Timeout: timeout,
	}
	requestsPerGoRoutine := numRequests / concurrency
	remainingRequests := numRequests % concurrency

	start := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendRequests(client, url, requestsPerGoRoutine)
		}()
	}

	if remainingRequests > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendRequests(client, url, remainingRequests)
		}()
	}

	wg.Wait()

	totalDuration = time.Since(start)

	fmt.Printf("Total requests: %d\n", numRequests)
	fmt.Printf("Concurrency level: %d\n", concurrency)
	fmt.Printf("Total duration: %v\n", totalDuration)

}

// Function to send HTTP requests to the URL
func sendRequests(client *http.Client, url string, count int) {
	for i := 0; i < count; i++ {
		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		resp.Body.Close()
	}

}
