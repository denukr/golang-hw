package main

import (
	"3-struct/app/bins"
	"3-struct/app/storage"
	"fmt"
)

func main() {
	myBin, _ := bins.NewBin("0", true, "myBin")

	myStorage, err := storage.NewStorage()
	if err != nil {
		fmt.Print(err)
	}
	myStorage.AddBin(*myBin)
	myStorage.Save()
	res, _ := storage.ReadBinList(myStorage)
	fmt.Println(res)
}
