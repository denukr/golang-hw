package file

import (
	"fmt"
	"os"
	"path/filepath"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(fileName string) *JsonDb {
	return &JsonDb{
		fileName: fileName,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	data, err := os.ReadFile(db.fileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func (db *JsonDb) Write(content []byte) {
	file, err := os.Create(db.fileName)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func checkJSON(path string) bool {
	fileExtension := filepath.Ext(path)
	if fileExtension == "json" {
		return true
	}
	return false
}
