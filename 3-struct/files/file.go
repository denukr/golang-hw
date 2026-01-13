package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type JsonDb struct {
	fileName string
}

func NewJsonDb(fileName string) (*JsonDb, error) {
	db := &JsonDb{
		fileName: fileName,
	}
	if !db.checkJSON() {
		return nil, errors.New("Incorrect_File_Extension")
	}
	return db, nil
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

func (db *JsonDb) checkJSON() bool {
	fileExtension := filepath.Ext(db.fileName)
	return fileExtension == ".json"
}
