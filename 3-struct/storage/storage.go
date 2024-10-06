package storage

import (
	"3-struct/app/bins"
	"encoding/json"
	"fmt"
	"time"
)

type Db interface {
	Read() ([]byte, error)
	Write(content []byte)
}

type Storage struct {
	Bins     []bins.Bin `json: "bins"`
	UpdateAt time.Time  `json: "updateAt"`
}

type StorageWithDb struct {
	Storage
	db Db
}

func (storage *StorageWithDb) AddBin(bin bins.Bin) {
	storage.Bins = append(storage.Bins, bin)
}

func (storage *StorageWithDb) Save() {
	storage.UpdateAt = time.Now()
	data, err := storage.ToBytes()
	if err != nil {
		fmt.Println(err)
		return
	}
	storage.db.Write(data)
}

func (storage *StorageWithDb) ToBytes() ([]byte, error) {
	data, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func NewStorage(db Db) (*StorageWithDb, error) {
	data, err := db.Read()
	if err != nil || len(data) == 0 {
		fmt.Println(err)
		return &StorageWithDb{
			Storage: Storage{},
			db:      db,
		}, nil
	}
	storage := Storage{}
	err = json.Unmarshal(data, &storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &StorageWithDb{
		Storage: storage,
		db:      db,
	}, nil
}
