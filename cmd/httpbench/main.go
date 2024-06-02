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
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
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

	color.Cyan("Starting benchmark for URL: %s\n", url)

	bar := pb.StartNew(numRequests)
	bar.SetWriter(os.Stderr)
	bar.SetRefreshRate(time.Millisecond * 100)
	bar.Set(pb.SIBytesPrefix, true)

	// Call benchmarking function
	m := benchmark.Run(url, numRequests, concurrency, timeout, bar)

	bar.Finish()

	color.Green("Benchmarking completed!\n")

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredBright)
	t.AppendHeader(table.Row{"Metric", "Value"})
	t.AppendRows([]table.Row{
		{"Total requests", m.TotalRequests},
		{"Successful requests", m.SuccessfulRequests},
		{"Failed requests", m.FailedRequests},
		{"Average response time", fmt.Sprintf("%.2f ms", m.AverageResponseTime().Seconds()*1000)},
		{"Throughput", fmt.Sprintf("%.2f req/s", m.RequestsPerSecond())},
	})

	t.Render()

	if m.FailedRequests > 0 {
		color.Red("Failed request detected!")
	} else {
		color.Green("All requests successful!")
	}
}
