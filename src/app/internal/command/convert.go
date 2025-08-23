package command

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

type Response struct {
	Success bool `json:"success"`
	Query   struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	} `json:"query"`
	Info struct {
		Rate float64 `json:"rate"`
	} `json:"info"`
	Result float64 `json:"result"`
}

var from, to string
var amount float64

var convertCmd = &cobra.Command{
	Use:   "",
	Short: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		url := fmt.Sprintf("https://api.exchangerate.host/convert?from=%s&to=%s&amount=%f&access_key=%s", from, to, amount, apiKey)

		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		var res Response
		if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
			return err
		}

		if !res.Success {
			return fmt.Errorf("asd")
		}

		fmt.Printf("%.2f %s = %.2f %s\n",
			res.Query.Amount, res.Query.From, res.Result, res.Query.To)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringVarP(&from, "from", "f", "USD", "source currency")
	convertCmd.Flags().StringVarP(&to, "to", "t", "EUR", "target currency")
	convertCmd.Flags().Float64VarP(&amount, "amount", "a", 1, "amount to  currency")
}
