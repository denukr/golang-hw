package main

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	usdEuro   = 0.91
	usdRuble  = 91.37
	euroRuble = usdRuble / usdEuro
)

func main() {
	fmt.Println("__Конверетр валют__")
	fmt.Println("Меню:")
	fmt.Println("1. Введите исходную валюту;")
	fmt.Println("2. Введите число;")
	fmt.Println("3. Введите целевую валюту.")
	fmt.Printf("\n")
	for {
		currentCurrancy, number, targetCurrency, err := getUserInput()
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
			continue
		}
		result, _ := convertCurrency(currentCurrancy, number, targetCurrency)
		fmt.Printf("Результат:%v %v=%.2f %v\n", number, currentCurrancy, result, targetCurrency)
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
		return "", number, "", err
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
	fmt.Print("1. Введите исходную валюту (USD, Euro или Ruble): ")
	fmt.Scan(&currentCurrency)
	if currentCurrency == "USD" || currentCurrency == "Euro" || currentCurrency == "Ruble" {
		return currentCurrency, nil
	} else {
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
	fmt.Printf("3. Введите целевую валюту (%v или %v): ", currencyA, currencyB)
	if fmt.Scan(&targetCurrency); targetCurrency == currencyA || targetCurrency == currencyB {
		return targetCurrency, nil
	}
	return "", errors.New("введена неверная целевая валюта")
}

func getNumber() (float64, error) {
	var numberStr string
	fmt.Print("3. Введите число: ")
	fmt.Scan(&numberStr)
	num, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, errors.New("введено не число")
	}
	number := float64(num)
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
	fmt.Print("Желаете еще конверировать валюту? (Y/N)")
	fmt.Scan(&userAnswer)
	fmt.Printf("\n")
	if userAnswer == "Y" || userAnswer == "y" {
		return true
	}
	return false
}
