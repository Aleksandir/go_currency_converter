package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ConversionResponse struct {
	Base       string             `json:"base"`
	Date       string             `json:"date"`
	Rates      map[string]float64 `json:"rates"`
}

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