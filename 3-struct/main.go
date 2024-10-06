package main

import (
	"3-struct/app/bins"
	file "3-struct/app/files"
	"3-struct/app/storage"
	"fmt"
)

func main() {
	myBin, _ := bins.NewBin("0", true, "myBin")

	myStorage, err := storage.NewStorage(file.NewJsonDb("data.json"))
	if err != nil {
		fmt.Print(err)
	}
	myStorage.AddBin(*myBin)
	myStorage.Save()
}
