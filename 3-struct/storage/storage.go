package storage

import (
	"3-struct/app/bins"
	file "3-struct/app/files"
	"encoding/json"
	"fmt"
	"time"
)

type Storage struct {
	Bins     []bins.Bin `json: "bins"`
	UpdateAt time.Time  `json: "updateAt"`
}

func (storage *Storage) AddBin(bin bins.Bin) {
	storage.Bins = append(storage.Bins, bin)
}

func (storage *Storage) Save() {
	storage.UpdateAt = time.Now()
	data, err := storage.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	file.WriteFile(data, "data.json")
}

func (storage *Storage) ToBytes() ([]byte, error) {
	data, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func NewStorage() (*Storage, error) {
	data, err := file.ReadFile("data.json")
	if err != nil || len(data) == 0 {
		fmt.Println(err)
		return &Storage{}, nil
	}
	storage := &Storage{}
	err = json.Unmarshal(data, storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return storage, nil
}

func ReadBinList(storage *Storage) ([]bins.Bin, error) {
	storage, err := NewStorage()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return storage.Bins, nil
}
