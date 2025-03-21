package cmd

import (
	"fmt"
	"os"

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
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("URL: %s\n", url)
		fmt.Printf("Requests: %d\n", requests)
		fmt.Printf("Concurrency: %d\n", concurrency)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado")
	rootCmd.Flags().IntVarP(&requests, "requests", "r", 1, "Número total de requests")
	rootCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Número de chamadas simultâneas")

	rootCmd.MarkFlagRequired("url")
}
