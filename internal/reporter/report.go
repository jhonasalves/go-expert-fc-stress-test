package reporter

import (
	"fmt"
	"time"

	"github.com/jhonasalves/go-expert-fc-stress-test/internal/tester"
)

func GenerateReport(results tester.TestResult, duration time.Duration) {
	fmt.Println("\n--- Load Test Report ---")
	fmt.Printf("Total execution time: %v\n", duration)
	fmt.Printf("Total requests made: %d\n", results.TotalRequests)
	fmt.Printf("Requests with status 200: %d\n", results.SuccessfulRequests)

	for statusCode, count := range results.StatusCodes {
		fmt.Printf("Status %d: %d requests\n", statusCode, count)
	}
}
