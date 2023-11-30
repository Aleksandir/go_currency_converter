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

	convertedAmount := converter.Convert(amt, rate)

	fmt.Printf("%v %v is %v %v\n", *amount, *fromCurrency, convertedAmount, *toCurrency)
}