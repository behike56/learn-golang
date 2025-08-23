package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var apiKey string

var rootCmd = &cobra.Command{
	Use:   "currency",
	Short: "Currency shor msg.",
	Long:  "Currency long msg.",
}

func Execute(key string) {

	if key == "" {
		fmt.Println("CURRENCY_API_KEY environment variable is not set.")
		os.Exit(1)
	}

	apiKey = key

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
