package entity

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	ID        uint64 `gorm:"primaryKey"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Book) TableName() string {
	return "books"
}
