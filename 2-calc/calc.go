package main

import (
	"fmt"
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
	var operation, nums string
	fmt.Print("Введите числа через запятую: ")
	fmt.Scan(&nums)
	fmt.Print("Введите операцию: ")
	fmt.Scan(&operation)
	a := convertStringToNumber(nums)
	menuFunc := menu[operation]
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
