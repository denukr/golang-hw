package main

import (
	"fmt"
	"strings"
)

type currencyMap = map[string]map[string]float64

func main() {
	currencyMap := currencyMap{
		"USD": map[string]float64{
			"Euro":  0.91,
			"Ruble": 91.37},
	}
	currencyMap["Euro"] = make(map[string]float64, 2)
	currencyMap["Euro"]["USD"] = 1 / currencyMap["USD"]["Euro"]
	currencyMap["Euro"]["Ruble"] = currencyMap["USD"]["Ruble"] / currencyMap["USD"]["Euro"]
	currencyMap["Ruble"] = make(map[string]float64, 2)
	currencyMap["Ruble"]["USD"] = 1 / currencyMap["USD"]["Ruble"]
	currencyMap["Ruble"]["Euro"] = 1 / currencyMap["Euro"]["Ruble"]
	fmt.Println("__Конверетр валют__")
	fmt.Println("Меню:")
	fmt.Println("1. Введите исходную валюту;")
	fmt.Println("2. Введите число;")
	fmt.Println("3. Введите целевую валюту.")
	fmt.Printf("\n")
	for {
		var currentCurrancy, targetCurrancy string
		var number float64
		fmt.Print("1. (Euro, USD или Ruble) ")
		fmt.Scan(&currentCurrancy)
		fmt.Print("2. ")
		fmt.Scan(&number)
		getTargetCurrency(&currentCurrancy, &targetCurrancy, currencyMap)
		result, _ := convertCurrency(currentCurrancy, number, targetCurrancy, currencyMap)
		fmt.Printf("Результат: %.2F %v\n", result, targetCurrancy)
		if !wantToContinue() {
			break
		}
	}
}

func convertCurrency(currentCurrancy string, number float64, targetCurrancy string, currencies currencyMap) (float64, error) {
	return currencies[currentCurrancy][targetCurrancy] * number, nil
}

func wantToContinue() bool {
	var userAnswer string
	fmt.Print("Желаете еще конверировать валюту? (Y/N) ")
	fmt.Scan(&userAnswer)
	if userAnswer == "Y" || userAnswer == "y" {
		return true
	}
	return false
}

func getTargetCurrency(currentCurrancy *string, targetCurrancy *string, currencyMap currencyMap) {
	str := ""
	for key := range currencyMap[*currentCurrancy] {
		str += key
		str += " или "
	}
	str = strings.TrimRight(str, "ил ")
	fmt.Printf("3. (%v) ", str)
	fmt.Scan(targetCurrancy)
}
