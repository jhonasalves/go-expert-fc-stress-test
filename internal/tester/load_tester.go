package tester

import (
	"fmt"
	"net/http"
	"sync"
)

type TestResult struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusCodes        map[int]int
}

func RunLoadTest(url string, totalRequests, concurrency int) TestResult {
	var wg sync.WaitGroup
	results := TestResult{
		TotalRequests:      totalRequests,
		SuccessfulRequests: 0,
		StatusCodes:        make(map[int]int),
	}

	// Channel for communication of results from goroutines
	resultChan := make(chan int, totalRequests) // Buffered channel with size equal to the total number of requests

	// Concurrency control semaphore: Limits the number of simultaneous goroutines
	sem := make(chan struct{}, concurrency) // Channel with "concurrency" slots

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{} // Occupies a slot in the semaphore to limit concurrency

		go func() {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error making request:", err)
				<-sem
				return
			}
			defer resp.Body.Close()

			resultChan <- resp.StatusCode

			<-sem // Releases a slot in the semaphore
		}()
	}

	// Waits for all goroutines to finish
	go func() {
		wg.Wait()
		close(resultChan) // Closes the channel only after all goroutines are done
	}()

	// Processes the status codes
	for status := range resultChan {
		if status == 200 {
			results.SuccessfulRequests++
		}
		results.StatusCodes[status]++
	}

	return results
}
