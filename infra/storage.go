package infra

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadJSON загружает данные из JSON-файла в переданный указатель на объект.
func LoadJSON[T any](path string, obj *T) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("\n%s not found, make new", path)
		return
	}

	err = json.Unmarshal(data, obj)
	if err != nil {
		fmt.Printf("\nLoading %s error: %s", path, err)
	}
}

// SaveJSON сохраняет объект в JSON-файл.
func SaveJSON[T any](path string, obj T) {
	os.MkdirAll(filepath.Dir(path), 0755)

	data, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		fmt.Printf("\nSaving %s error: %s", path, err.Error())
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Printf("\nWriting %s error: %s", path, err.Error())
	}
}
