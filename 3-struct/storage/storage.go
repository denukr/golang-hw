package storage

import (
	"3-struct/app/bins"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type BinInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Db interface {
	Read(id string) ([]byte, error)
	Write(data []byte, fileName string)
	Update(id string, data []byte)
	Delete(id string) error
	List() ([]BinInfo, error)
}

type Storage struct {
	Bins     []bins.Bin `json:"bins"`
	UpdateAt time.Time  `json:"updateAt"`
}

type StorageWithDb struct {
	Storage
	db Db
}

// SaveNew теперь просто пересылает файл в DB (API)
func (s *StorageWithDb) SaveNew(localPath string, cloudName string) {
	// Читаем сырые байты из файла (например, {"Hello": 1})
	content, err := os.ReadFile(localPath)
	if err != nil {
		fmt.Printf("Ошибка: файл %s не найден\n", localPath)
		return
	}

	// Отправляем эти байты в API
	s.db.Write(content, cloudName)
}

func (s *StorageWithDb) UpdateRemote(id string) {
	s.UpdateAt = time.Now()
	data, _ := json.MarshalIndent(s, "", "  ")
	s.db.Update(id, data)
}

func NewStorage(db Db, id string) (*StorageWithDb, error) {
	data, err := db.Read(id)
	if err != nil || len(data) == 0 {
		return &StorageWithDb{Storage: Storage{}, db: db}, nil
	}
	var s Storage
	err = json.Unmarshal(data, &s)
	return &StorageWithDb{Storage: s, db: db}, err
}
