package tester

import (
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/briandowns/spinner"
)

type TestResult struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusCodes        sync.Map
}

func RunLoadTest(url string, totalRequests, concurrency int) *TestResult {
	var wg sync.WaitGroup
	var successfulRequests int32

	results := &TestResult{
		TotalRequests: totalRequests,
	}

	// Channel for communication of results from goroutine
	resultChan := make(chan int, totalRequests)
	// Concurrency control semaphore: Limits the number of simultaneous goroutines
	sem := make(chan struct{}, concurrency)

	s := spinner.New(spinner.CharSets[36], 300*time.Millisecond)
	s.Start()
	defer s.Stop()

	transport := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: concurrency,
	}
	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer func() {
				wg.Done()
				<-sem
			}()

			// Releases the slot in the semaphore
			resp, err := client.Get(url)
			if err != nil {
				resultChan <- 0 // Indicates an error without printing to the terminal
				return
			}
			defer resp.Body.Close()

			resultChan <- resp.StatusCode
		}()
	}

	// Waits for all goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Processes the status codes
	for status := range resultChan {
		if status == 200 {
			atomic.AddInt32(&successfulRequests, 1)
		}

		if count, ok := results.StatusCodes.Load(status); ok {
			results.StatusCodes.Store(status, count.(int)+1)
		} else {
			results.StatusCodes.Store(status, 1)
		}
	}

	results.SuccessfulRequests = int(successfulRequests)
	return results
}
