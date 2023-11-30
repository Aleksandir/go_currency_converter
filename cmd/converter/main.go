package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"currency-converter/pkg/api"
	"currency-converter/pkg/converter"
)

func main() {
	fromCurrency := flag.String("from", "USD", "Currency to convert from")
	toCurrency := flag.String("to", "EUR", "Currency to convert to")
	amount := flag.String("amount", "1", "Amount to convert")

	flag.Parse()

	amt, err := strconv.ParseFloat(*amount, 64)
	if err != nil {
		fmt.Println("Invalid amount. Please enter a valid number.")
		os.Exit(1)
	}

	rate, err := api.GetConversionRate(*fromCurrency, *toCurrency)
	if err != nil {
		fmt.Println("Error getting conversion rate:", err)
		os.Exit(1)
	}

	if rate <= 0 {
		fmt.Println("Received invalid conversion rate:", rate)
		os.Exit(1)
	}

	convertedAmount, err := converter.Convert(amt, *fromCurrency, *toCurrency)
	if err != nil {
		fmt.Println("Error converting amount:", err)
		os.Exit(1)
	}

	fmt.Printf("%v %v is %v %v\n", *amount, *fromCurrency, convertedAmount, *toCurrency)
}
