package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var menu = map[string]func([]float64) float64{
	"SUM": calcSUM,
	"AVG": calcAVG,
	"MED": calcMED,
}

func main() {
	nums, err := promptData("Введите числа через запятую")
	if err != nil {
		return
	}
	fmt.Println(nums)
	operation, err := promptData("Введите операцию")
	if err != nil {
		return
	}
	a := convertStringToNumber(nums)
	menuFunc := menu[operation]
	if menuFunc == nil {
		return
	}
	fmt.Println(menuFunc(a))
}

func convertStringToNumber(nums string) []float64 {
	splitNums := strings.Split(nums, ",")
	numArray := []float64{}
	for i := range splitNums {
		splitNums[i] = strings.TrimLeft(splitNums[i], " ")
		aFloat, _ := strconv.ParseFloat(splitNums[i], 64)
		numArray = append(numArray, aFloat)
	}
	return numArray
}

func calcAVG(nums []float64) float64 {
	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}
	return sum / float64(len(nums))
}

func calcSUM(nums []float64) float64 {
	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func calcMED(nums []float64) float64 {
	numsCopy := nums
	sort.Float64s(numsCopy)

	if len(numsCopy)%2 != 0 {
		return numsCopy[(len(numsCopy)-1)/2]
	}
	return calcAVG(numsCopy[(len(numsCopy)-1)/2 : (len(numsCopy)-1)/2+2])

}

func promptData(prompt ...any) (string, error) {
	for i, val := range prompt {
		if i == len(prompt)-1 {
			fmt.Printf("%v: ", val)
		} else {
			fmt.Println(val)
		}
	}
	var result string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		result = scanner.Text()
	}
	err := scanner.Err()
	if err != nil {
		return "", errors.New("Stdin_ERROR")
	}
	return result, nil
}
