package main

import "fmt"

func main() {
	const (
		usdEuro = 0.91
		usdRuble = 91.37
		euroRuble = usdRuble / usdEuro
	)
	fmt.Printf("%v %v %.2f", usdEuro, usdRuble, euroRuble)
}