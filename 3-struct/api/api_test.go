package api_test

import (
	"3-struct/app/api"
	"encoding/json"
	"fmt" // Используем для мгновенного вывода
	"testing"
)

func TestWrite(t *testing.T) {
	fmt.Println("--- СТАРТ: TestWrite ---")

	type TestData struct {
		Text string `json:"text"`
	}

	data := TestData{Text: "Hello"}
	testData, _ := json.Marshal(data)

	db := &api.RemoteDb{URL: "https://api.jsonbin.io"}

	// Act
	fmt.Printf("Отправка данных в API: %s\n", string(testData))
	binId, err := db.Write(testData, "test")
	if err != nil {
		t.Fatalf("Ошибка записи: %v", err)
	}

	// Гарантируем удаление после теста
	t.Cleanup(func() {
		db.Delete(binId)
		fmt.Printf("Бин %s удален\n", binId)
	})

	// Assert
	got, err := db.Read(binId)
	if err != nil {
		t.Fatalf("Ошибка чтения: %v", err)
	}

	fmt.Printf("Получено из API: %s\n", string(got))

	if string(got) != string(testData) {
		t.Errorf("ОЖИДАЛИ %s, ПОЛУЧИЛИ %s", string(testData), string(got))
	}
}

func TestUpdate(t *testing.T) {
	fmt.Println("--- СТАРТ: TestUpdate ---")

	db := &api.RemoteDb{URL: "https://api.jsonbin.io"}

	// Подготовка
	initialData, _ := json.Marshal(map[string]string{"text": "Hello"})
	updateData, _ := json.Marshal(map[string]string{"text": "HelloUpdated"})

	binId, err := db.Write(initialData, "test")
	if err != nil {
		t.Fatalf("Не удалось создать бин: %v", err)
	}
	t.Cleanup(func() { db.Delete(binId) })

	// Act
	fmt.Println("Обновление данных...")
	err = db.Update(binId, updateData)
	if err != nil {
		t.Fatalf("Ошибка при обновлении: %v", err)
	}

	// Assert
	got, _ := db.Read(binId)
	fmt.Printf("Проверка после обновления: %s\n", string(got))

	if string(got) != string(updateData) {
		t.Errorf("Данные не обновились. Получено: %s", string(got))
	}
}

func TestRead(t *testing.T) {
	fmt.Println("--- СТАРТ: TestRead ---")

	db := &api.RemoteDb{URL: "https://api.jsonbin.io"}

	// 1. Подготовка: сначала записываем данные, которые будем читать
	originalData := map[string]string{"status": "active", "msg": "read_test"}
	payload, _ := json.Marshal(originalData)

	fmt.Println("Шаг 1: Создание временного бина...")
	binId, err := db.Write(payload, "test_read")
	if err != nil {
		t.Fatalf("Не удалось создать данные для теста: %v", err)
	}

	// Очистка: удаляем бин после теста
	t.Cleanup(func() {
		fmt.Printf("Очистка: удаление бина %s\n", binId)
		db.Delete(binId)
	})

	// 2. Действие: чтение созданного бина
	fmt.Printf("Шаг 2: Запрос на чтение бина %s...\n", binId)
	got, err := db.Read(binId)
	if err != nil {
		t.Fatalf("Ошибка при чтении: %v", err)
	}

	// 3. Проверка
	fmt.Printf("Шаг 3: Сверка данных. Получено: %s\n", string(got))

	var receivedData map[string]string
	if err := json.Unmarshal(got, &receivedData); err != nil {
		t.Fatalf("Не удалось распарсить ответ API: %v", err)
	}

	if receivedData["msg"] != originalData["msg"] {
		t.Errorf("Данные не совпадают! Ожидали %s, получили %s", originalData["msg"], receivedData["msg"])
	}

	fmt.Println("--- УСПЕХ: TestRead пройден ---")
}

func TestDelete(t *testing.T) {

	fmt.Println("--- TestDelete ---")
	// Arrange
	originalData := map[string]string{"status": "active", "msg": "read_test"}
	payload, err := json.Marshal(originalData)
	if err != nil {
		t.Fatalf("Не удалось создать данные для теста %v", err)
	}

	db := api.RemoteDb{URL: "https://api.jsonbin.io"}
	binId, err := db.Write(payload, "test_delete")
	if err != nil {
		t.Fatalf("Не удалось создать бин")
	}

	// Act
	err = db.Delete(binId)
	if err != nil {
		t.Fatalf("Не удалось удалить бин с ID=%v", binId)
	}

	// Assert
	_, err = db.Read(binId)
	if err == nil {
		t.Fatalf("Прочитали бин с id=%v, хотя ожидалась ошибка", binId)
	}

}
