package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ExchangeRates struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func (e *ExchangeRates) Convert(amount float64, currency string) (float64, error) {
	rate, ok := e.Rates[strings.ToUpper(currency)]
	if !ok {
		return 0, fmt.Errorf("currency %s not found in the exchange rates list", currency)
	}
	return amount * rate, nil
}

func swapDecimalPointForComma(amount float64) string {
	s := fmt.Sprintf("%.2f", amount)
	return strings.Replace(s, ".", ",", 1)
}

func main() {
	argsWithoutProgram := os.Args[1:]
	if len(argsWithoutProgram) != 2 {
		fmt.Fprintln(os.Stderr, "You should pass 2 arguments: Amount as a decimal and currency as a three letter code")
		os.Exit(1)
	}

	amount, err := strconv.ParseFloat(argsWithoutProgram[0], 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Amount is invalid. Should be in the following format: 3.14")
		os.Exit(1)
	}

	currency := argsWithoutProgram[1]
	if len(currency) != 3 {
		fmt.Fprintln(os.Stderr, "Currency is in the wrong format. Should be a currency code. Example: BRL, EUR, USD")
		os.Exit(1)
	}

	fileData, err := os.ReadFile("./rates.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not read rates file:", err)
		os.Exit(1)
	}

	exchangeRates := &ExchangeRates{}

	err = json.Unmarshal(fileData, exchangeRates)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error while parsing the exchange rates json:", err)
		os.Exit(1)
	}

	convertedAmount, err := exchangeRates.Convert(amount, currency)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", swapDecimalPointForComma(convertedAmount))
}
