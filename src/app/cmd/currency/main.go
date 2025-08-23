package main

import (
	"os"

	"github.com/behike56/learn-golang/app/internal/command"
)

func main() {
	apiKey := os.Getenv("CURRENCY_API_KEY")
	command.Execute(apiKey)
}
