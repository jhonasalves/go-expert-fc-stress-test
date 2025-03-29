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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
