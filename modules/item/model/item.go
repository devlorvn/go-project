package model

import (
	"errors"

	"github.com/devlorvn/go-project/common"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleIsBlank  = errors.New("title cannot be blank")
	ErrItemIsDeleted = errors.New("Item has been deleted")
)

type TodoItem struct {
	common.BaseModel
	Title string `json:"title" gorm:"column:title;"`
	// Image
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItem) TableName() string { return "todo_items" }

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string     `json:"title" gorm:"column:title;"`
	Description *string     `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
