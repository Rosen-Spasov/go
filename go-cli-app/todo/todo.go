package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	Text string
}

func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, _ := os.ReadFile(filename)
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}
	return items, nil
}
