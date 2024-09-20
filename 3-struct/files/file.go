package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
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
