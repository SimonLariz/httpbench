// Define data structures for storing and calculating benchmark results
// Create functions for calculating metrics like latency, throughput, and error rate
// Create functions for aggregating and printing benchmark results
package metrics

import "time"

type Metrics struct {
	TotalRequests      int
	SuccessfulRequests int
	FailedRequests     int
	TotalDuration      time.Duration
	RequestDurations   []time.Duration
}

func NewMetrics(TotalRequests int) *Metrics {
	return &Metrics{
		TotalRequests:    TotalRequests,
		RequestDurations: make([]time.Duration, 0, TotalRequests),
	}
}

func (m *Metrics) AddRequestDuration(duration time.Duration) {
	m.RequestDurations = append(m.RequestDurations, duration)
}

func (m *Metrics) AddSuccessfulRequest() {
	m.SuccessfulRequests++
}

func (m *Metrics) AddFailedRequest() {
	m.FailedRequests++
}

func (m *Metrics) SetTotalDuration(duration time.Duration) {
	m.TotalDuration = duration
}

func (m *Metrics) AverageResponseTime() time.Duration {
	if len(m.RequestDurations) == 0 {
		return 0
	}
	var sum time.Duration
	for _, duration := range m.RequestDurations {
		sum += duration
	}
	return sum / time.Duration(len(m.RequestDurations))
}

func (m *Metrics) RequestsPerSecond() float64 {
	if m.TotalDuration == 0 {
		return 0
	}
	return float64(m.TotalRequests) / m.TotalDuration.Seconds()
}
