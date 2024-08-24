package main

import "fmt"

func main() {
	const (
		usdEuro   = 0.91
		usdRuble  = 91.37
		euroRuble = usdRuble / usdEuro
	)
	getUserInput()
	convertCurrency(4, "", "")
}

func getUserInput() {
	// var currentCurrency, targetCurrency float64
	var someString string
	fmt.Scan(&someString)
	fmt.Print(someString)
}

func convertCurrency(c int, a string, b string) {}
