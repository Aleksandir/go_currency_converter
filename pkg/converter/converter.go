package converter

import (
	"currency-converter/pkg/api"
	"errors"
)

// Convert function takes an amount and two currency codes as arguments and returns the converted amount.
func Convert(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	if amount <= 0 {
		return 0, errors.New("amount must be greater than zero")
	}

	rate, err := api.GetConversionRate(fromCurrency, toCurrency)
	if err != nil {
		return 0, err
	}

	return amount * rate, nil
}
