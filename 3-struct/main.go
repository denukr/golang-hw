package main

import (
	"fmt"

	"3-struct/app/api"
	"3-struct/app/bins"
)

func main() {
	testBin, _ := bins.NewBin("1", true, "test")
	s := *testBin
	fmt.Println(s)
	api.Api()
}
