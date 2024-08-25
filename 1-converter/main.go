package main

import (
	"errors"
	"fmt"
)

const (
	usdEuro   = 0.91
	usdRuble  = 91.37
	euroRuble = usdRuble / usdEuro
)

func main() {
	fmt.Println("__Конверетр валют__")
	for {
		currentCurrancy, number, targetCurrency, err := getUserInput()
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, _ := convertCurrency(currentCurrancy, number, targetCurrency)
		fmt.Printf("Результат: %v\n", result)
		if !wantToContinue() {
			break
		}
	}
}

func getUserInput() (string, float64, string, error) {
	currentCurrancy, err := getCurrentCurrency()
	for err != nil {
		return "", 0, "", err
	}
	number, err := getNumber()
	if err != nil {
		fmt.Printf("Что за чило?")
		return "", 0, "", err
	}
	curA, curB, err := returnRemainingCurrencies(currentCurrancy)
	if err != nil {
		return "", 0, "", err
	}
	targetCurrency, err := getTargetCurrency(curA, curB)
	if err != nil {
		return "", 0, "", err
	}
	return currentCurrancy, number, targetCurrency, nil
}

func getCurrentCurrency() (string, error) {
	var currentCurrency string
	fmt.Print("Введите исходную валюту (USD, Euro или Ruble): ")
	switch fmt.Scan(&currentCurrency); currentCurrency {
	case "USD":
		return currentCurrency, nil
	case "Euro":
		return currentCurrency, nil
	case "Ruble":
		return currentCurrency, nil
	default:
		return "", errors.New("вы пытаетесь ввести валюту, которой у нас нет")
	}
}

func returnRemainingCurrencies(currentCurrency string) (string, string, error) {
	switch currentCurrency {
	case "USD":
		return "Euro", "Ruble", nil
	case "Euro":
		return "USD", "Ruble", nil
	case "Ruble":
		return "USD", "Euro", nil
	default:
		return "", "", errors.New("получена неверная валюта для сравнения")
	}
}

func getTargetCurrency(currencyA string, currencyB string) (string, error) {
	var targetCurrency string
	fmt.Printf("Введите целевую валюту (%v или %v): ", currencyA, currencyB)
	if fmt.Scan(&targetCurrency); targetCurrency == currencyA || targetCurrency == currencyB {
		return targetCurrency, nil
	}
	return "", errors.New("введена неверная целевая валюта")
}

func getNumber() (float64, error) {
	var number float64
	fmt.Print("Введите число: ")
	fmt.Scan(&number)
	if number <= 0 {

		return 0, errors.New("введено некорректное число")
	}
	return number, nil
}

func convertCurrency(currentCurrancy string, number float64, targetCurrancy string) (float64, error) {
	switch currentCurrancy {
	case "USD":
		if targetCurrancy == "Euro" {
			return number * usdEuro, nil
		} else if targetCurrancy == "Ruble" {
			return number * usdRuble, nil
		}
	case "Euro":
		if targetCurrancy == "USD" {
			return number / usdEuro, nil
		} else if targetCurrancy == "Ruble" {
			return number * euroRuble, nil
		}
	case "Ruble":
		if targetCurrancy == "USD" {
			return number / usdRuble, nil
		} else if targetCurrancy == "Euro" {
			return number / euroRuble, nil
		}
	}
	return 0, errors.New("пытаетесь конверитировать несуществуюую валюту")
}

func wantToContinue() bool {
	var userAnswer string
	fmt.Print("Желаете продолжить? (Y/N)")
	fmt.Scan(&userAnswer)
	if userAnswer == "Y" || userAnswer == "y" {
		return true
	}
	return false
}
