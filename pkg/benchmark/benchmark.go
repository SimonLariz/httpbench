// Core benchmarking functionality
// Function takes URL and sends HTTP requests to the URL using net/http package
// Implement logic for response time calculation
// Use goroutines to send multiple requests concurrently
package benchmark

import (
	"net/http"
	"sync"
	"time"

	"github.com/SimonLariz/httpbench/pkg/metrics"
)

func Run(url string, numRequests, concurrency int, timeout time.Duration) *metrics.Metrics {
	// Implement benchmarking logic
	var wg sync.WaitGroup
	m := metrics.NewMetrics(numRequests)

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
			sendRequests(client, url, requestsPerGoRoutine, m)
		}()
	}

	if remainingRequests > 0 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sendRequests(client, url, remainingRequests, m)
		}()
	}

	wg.Wait()

	totalDuration := time.Since(start)
	m.SetTotalDuration(totalDuration)

	return m
}

// Function to send HTTP requests to the URL
func sendRequests(client *http.Client, url string, count int, m *metrics.Metrics) {
	for i := 0; i < count; i++ {
		start := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(start)
		m.AddRequestDuration(duration)

		if err != nil {
			// Uncomment the line below to print the error message
			// fmt.Println("Error:", err)
			m.AddFailedRequest()
			continue
		}
		resp.Body.Close()
		m.AddSuccessfulRequest()
	}
}
