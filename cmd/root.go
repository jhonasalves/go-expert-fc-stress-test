package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jhonasalves/go-expert-fc-stress-test/internal/reporter"
	"github.com/jhonasalves/go-expert-fc-stress-test/internal/tester"
	"github.com/spf13/cobra"
)

var (
	url         string
	requests    int
	concurrency int
)

var rootCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "Load testing tool for HTTP services",
	Long: `LoadTest is a CLI tool designed to stress test HTTP services.
It allows you to simulate multiple concurrent requests to evaluate the performance
and reliability of your application under heavy load.

Features:
- Specify the target URL to test.
- Configure the total number of requests to send.
- Set the level of concurrency for simultaneous requests.

Example usage:
  loadtest --url https://example.com --requests 1000 --concurrency 50
`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if url == "" {
			log.Fatal("URL is required. Use --url or -u to provide a URL.")
		}

		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			log.Fatal("Invalid URL. URL should start with 'http://' or 'https://'.")
		}

		if requests <= 0 {
			log.Fatal("Invalid number of requests. It should be greater than 0.")
		}

		if concurrency <= 0 {
			log.Fatal("Invalid concurrency. It should be greater than 0.")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Requests: %d\n", requests)
		fmt.Printf("Concurrency: %d\n", concurrency)

		start := time.Now()
		results := tester.RunLoadTest(url, requests, concurrency)
		duration := time.Since(start)

		reporter.GenerateReport(results, duration)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL of the service to be tested")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 1, "Total number of requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Number of simultaneous calls")

	rootCmd.MarkFlagRequired("url")
}
