package main

import (
	"3-struct/app/api"
	"3-struct/app/storage"
	"flag"
	"fmt"
)

func main() {
	// Флаги аргументов
	filePath := flag.String("file", "", "Путь к локальному файлу (откуда берем данные)")
	binName := flag.String("name", "", "Имя бина в облаке (X-Bin-Name)")
	id := flag.String("id", "", "ID бина для работы в облаке")

	// Флаги-операции
	createOp := flag.Bool("create", false, "Создать новый бин")
	updateOp := flag.Bool("update", false, "Обновить существующий бин")
	deleteOp := flag.Bool("delete", false, "Удалить бин")
	getOp := flag.Bool("get", false, "Получить данные бина")
	listOp := flag.Bool("list", false, "Показать список")

	flag.Parse()

	db := &api.RemoteDb{URL: "https://api.jsonbin.io"}

	if *listOp {
		items, _ := db.List()
		fmt.Println("ID | Name")
		for _, item := range items {
			fmt.Printf("%s | %s\n", item.ID, item.Name)
		}
		return
	}

	s, _ := storage.NewStorage(db, "")

	if *createOp {
		if *filePath == "" {
			fmt.Println("Ошибка: укажите --file")
			return
		}

		cloudName := *binName
		if cloudName == "" {
			cloudName = *filePath // Если --name нет, берем имя файла
		}

		// Вызываем метод, который просто перешлет байты
		s.SaveNew(*filePath, cloudName)
		return
	}

	if *getOp {
		if *id == "" {
			fmt.Println("Ошибка: нужен --id")
			return
		}
		data, _ := db.Read(*id)
		fmt.Println(string(data))
		return
	}

	if *updateOp {
		if *id == "" {
			fmt.Println("Ошибка: нужен --id")
			return
		}
		s.UpdateRemote(*id)
		return
	}

	if *deleteOp {
		if *id == "" {
			fmt.Println("Ошибка: нужен --id")
			return
		}
		_ = db.Delete(*id)
		fmt.Println("Удалено")
		return
	}
}
