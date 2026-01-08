package api

import (
	"3-struct/app/storage"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type RemoteDb struct {
	URL string
}

func (db *RemoteDb) Write(data []byte, fileName string) (binId string, err error) {
	req, _ := http.NewRequest(http.MethodPost, db.URL+"/v3/b", bytes.NewBuffer(data))

	// Это имя отобразится в панели управления JSONBin
	req.Header.Add("X-Bin-Name", fileName)
	req.Header.Add("X-Master-Key", "$2a$10$u88mSOCNGh1sLLDIJ4YNluW7rrhNUCTAxPWswUGXSOTqp.nepc8Xm")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка сети")
		return "", errors.New("ErrorNetwork")
	}
	defer resp.Body.Close()

	var res struct {
		Metadata struct {
			ID string `json:"id"`
		} `json:"metadata"`
	}
	body, _ := io.ReadAll(resp.Body)
	json.Unmarshal(body, &res)

	// Записываем в локальный список, чтобы потом работал --list
	db.updateLocalList(res.Metadata.ID, fileName)
	fmt.Printf("Успешно! Бин создан. ID: %s\n", res.Metadata.ID)
	return res.Metadata.ID, nil
}

func (db *RemoteDb) updateLocalList(id, name string) {
	var list []storage.BinInfo
	data, _ := os.ReadFile("list.json")
	json.Unmarshal(data, &list)
	list = append(list, storage.BinInfo{ID: id, Name: name})
	newData, _ := json.MarshalIndent(list, "", "  ")
	os.WriteFile("list.json", newData, 0644)
}

// Read получает содержимое бина (GET). Обрабатывает обертку "record" от JSONBin
func (db *RemoteDb) Read(id string) ([]byte, error) {
	if id == "" {
		return nil, nil
	}

	// 1. Создаем запрос вручную через NewRequest, чтобы иметь возможность добавить заголовки
	fullURL := fmt.Sprintf("%s/v3/b/%s/latest", db.URL, id)
	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return nil, err
	}

	// 2. Добавляем ключ авторизации (тот же, что вы используете в Write)
	req.Header.Add("X-Master-Key", "$2a$10$u88mSOCNGh1sLLDIJ4YNluW7rrhNUCTAxPWswUGXSOTqp.nepc8Xm")

	// 3. Выполняем запрос через клиент
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 4. Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("ошибка API (статус %d): %s", resp.StatusCode, string(body))
	}

	// 5. Читаем тело ответа
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 6. JSONBin возвращает данные в обертке {"record": ...}
	// Нужно извлечь само содержимое
	var wrapper struct {
		Record json.RawMessage `json:"record"`
	}
	if err := json.Unmarshal(body, &wrapper); err == nil && wrapper.Record != nil {
		return wrapper.Record, nil
	}

	return body, nil
}

func (db *RemoteDb) Update(id string, data []byte) error {
	req, _ := http.NewRequest(http.MethodPut, db.URL+"/v3/b/"+id, bytes.NewBuffer(data))
	req.Header.Add("X-Master-Key", "$2a$10$u88mSOCNGh1sLLDIJ4YNluW7rrhNUCTAxPWswUGXSOTqp.nepc8Xm")
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err == nil {
		resp.Body.Close()
		fmt.Println("Bin updated successfully")
		return nil
	} else {
		return err
	}
}

func (db *RemoteDb) Delete(id string) error {
	req, _ := http.NewRequest(http.MethodDelete, db.URL+"/v3/b/"+id, nil)
	req.Header.Add("X-Master-Key", "$2a$10$u88mSOCNGh1sLLDIJ4YNluW7rrhNUCTAxPWswUGXSOTqp.nepc8Xm")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (db *RemoteDb) List() ([]storage.BinInfo, error) {
	data, err := os.ReadFile("list.json")
	if err != nil {
		return nil, err
	}
	var list []storage.BinInfo
	err = json.Unmarshal(data, &list)
	return list, err
}
