package repository

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
)

// BookRepository 负责书籍相关的数据库操作
type BookRepository struct{}

// FindBooksByName 根据书名查询符合书籍
func (r *BookRepository) FindBooksByName(name string) ([]models.Book, error) {
	var books []models.Book
	err := utils.DB.Where("name LIKE ?", "%"+name+"%").Find(&books).Error
	return books, err
}

// FindBookByID 根据 ID 查询单个书籍信息
func (r *BookRepository) FindBookByID(bookID string) (models.Book, error) {
	var book models.Book
	err := utils.DB.First(&book, bookID).Error
	return book, err
}
