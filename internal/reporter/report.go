package reporter

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/jhonasalves/go-expert-fc-stress-test/internal/tester"
)

func GenerateReport(results tester.TestResult, duration time.Duration) {
	successColor := color.New(color.FgGreen).PrintfFunc()
	errorColor := color.New(color.FgRed).PrintfFunc()
	infoColor := color.New(color.FgCyan).PrintfFunc()

	successColor("\n--- Load Test Report ---\n")
	infoColor("Total execution time: %v\n", duration)
	infoColor("Total requests made: %d\n", results.TotalRequests)
	infoColor("Requests with status 200: %d\n\n", results.SuccessfulRequests)

	fmt.Println("\n---------------------- Status Codes ----------------------")

	fmt.Printf("%-10s %-20s\n", "Status", "Number of Requests")
	fmt.Println("---------------------------------------------------------")
	for statusCode, count := range results.StatusCodes {
		if statusCode == 200 {
			successColor("%-10d %-20d\n", statusCode, count)
		} else {
			errorColor("%-10d %-20d\n", statusCode, count)
		}
	}

	successColor("\nTest completed successfully!\n")
}
