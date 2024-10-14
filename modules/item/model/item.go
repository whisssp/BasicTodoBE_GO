package model

import (
	"ToDoAppBackEnd/common"
	"errors"
)

const (
	EntityName = "Item"
)

var (
	ErrTitleIsBlank = errors.New("title can not be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type TodoItem struct {
	common.SQLModel
	Title       string      `json:"title"  gorm:"column:title;"`
	Status      *ItemStatus `json:"status"  gorm:"column:status;"`
	Description string      `json:"description"  gorm:"column:description;"`
}

type TodoItemCreation struct {
	Id          int         `json:"-" gorm:"column:id;"`
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
}

type TodoItemUpdate struct {
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
	DD          string      `json:"test" gorm:"-"`
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }

func (TodoItem) TableName() string { return "todo_items" }
