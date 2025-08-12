package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	// Image
	Description string     `json:"description"`
	Status      string     `json:"status"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

func main() {
	fmt.Println("Hello world!!!!")

	now := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Item 1",
		Description: "Desc 1",
		Status:      "Doing",
		CreatedAt:   &now,
		UpdatedAt:   nil,
	}

	jsonData, err := json.Marshal(item)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	jsonStr := `{"id":1,"title":"Item 1","description":"Desc 1","status":"Doing","created_at":"2025-08-12T06:46:25.1694938Z","updated_at":null}`

	var item2 TodoItem

	if err := json.Unmarshal([]byte(jsonStr), &item2); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item2)
}
