package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConversionResponse struct {
	// Base is the base currency for the conversion.
	Base string `json:"base"`
	// Date is the date of the conversion rates.
	Date string `json:"date"`
	// Rates is a map of currency codes to conversion rates.
	Rates map[string]float64 `json:"rates"`
}

// GetConversionRate retrieves the conversion rate between two currencies.
// It makes a GET request to the exchangerate-api.com API to fetch the latest rates.
// The fromCurrency parameter specifies the currency to convert from.
// The toCurrency parameter specifies the currency to convert to.
// It returns the conversion rate as a float64 and an error if any.
func GetConversionRate(fromCurrency string, toCurrency string) (float64, error) {
	url := fmt.Sprintf("https://api.exchangerate-api.com/v4/latest/%s", fromCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var conversionResponse ConversionResponse
	err = json.Unmarshal(body, &conversionResponse)
	if err != nil {
		return 0, err
	}

	rate, exists := conversionResponse.Rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("conversion rate for '%s' to '%s' not found", fromCurrency, toCurrency)
	}

	return rate, nil
}
