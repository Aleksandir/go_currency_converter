package main

import (
	"fmt"
	"os"
	"strconv"

	"currency-converter/pkg/api"
	"currency-converter/pkg/converter"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: converter <from_currency> <to_currency> <amount>")
		os.Exit(1)
	}

	fromCurrency := os.Args[1]
	toCurrency := os.Args[2]
	amountStr := os.Args[3]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount. Please enter a valid number.")
		os.Exit(1)
	}

	rate, err := api.GetConversionRate(fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Error getting conversion rate:", err)
		os.Exit(1)
	}

	if rate <= 0 {
		fmt.Println("Received invalid conversion rate:", rate)
		os.Exit(1)
	}

	convertedAmount, err := converter.Convert(amount, fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Error converting amount:", err)
		os.Exit(1)
	}

	fmt.Printf("%v %v is %v %v\n", amount, fromCurrency, convertedAmount, toCurrency)
}
