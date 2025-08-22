package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func LoadJSON[T any](filename string) []T {
	data, err := os.ReadFile(filename)
	if err != nil {
		return []T{}
	}
	var items []T
	json.Unmarshal(data, &items)
	return items
}

func SaveJSON[T any](filename string, items []T) {
	data, _ := json.MarshalIndent(items, "", "  ")
	ioutil.WriteFile(filename, data, 0644)
}
