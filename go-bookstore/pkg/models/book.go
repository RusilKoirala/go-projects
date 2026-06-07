package models

import (
	"github.com/rusilkoirala/go-projects/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name"json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func CreateBook(db *gorm.DB, b *Book) error {
	return db.Create(b).Error
}

func DeleteBook(db *gorm.DB, id string) error {
	return db.Delete(&Book{}, id).Error
}

func UpdateBook(db *gorm.DB, b *Book, id string) error {
	return db.Model(&Book{}).
		Where("id = ?", id).
		Updates(b).
		Error
}

func GetAllBooks(db *gorm.DB) []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookbyId(db *gorm.DB, id string) (Book, error) {
	var book Book
	err := db.Where("id = ?", id).First(&book).Error
	if err != nil {
		return Book{}, err
	}
	return book, nil
}
