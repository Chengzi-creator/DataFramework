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
	if err != nil {
		return nil, err
	}
	return books, nil
}

// FindBookByID 根据 ID 查询单个书籍信息
func (r *BookRepository) FindBookByID(bookID int) (*models.Book, error) {
	var book models.Book
	err := utils.DB.First(&book, bookID).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// FindBookByName 根据书名查询书籍，用于缺书申请查找外键
func (r *BookRepository) FindBookByName(name string) (*models.Book, error) {
	var book models.Book
	err := utils.DB.Where("name=?", name).Error
	if err != nil {
		return nil, err
	}
	return &book, nil
}

// CreateBook 新增书籍
func (r *BookRepository) CreateBook(book *models.Book) error {
	err := utils.DB.Create(book).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateBook 更新书籍信息
func (r *BookRepository) UpdateBook(book *models.Book) error {
	err := utils.DB.Save(book).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteBookByBookID 根据bookID删除书籍
func (r *BookRepository) DeleteBookByBookID(bookID int) error {
	err := utils.DB.Delete(&models.Book{}, bookID).Error
	if err != nil {
		return err
	}
	return nil
}

// ShowBookShortage 显示缺书登记
func (r *BookRepository) ShowBookShortage() ([]models.BookShortage, error) {
	var bookShortageList []models.BookShortage
	err := utils.DB.Find(&bookShortageList).Error
	if err != nil {
		return nil, err
	}
	return bookShortageList, nil
}

// CreateBookShortage 创建缺书记录
func (r *BookRepository) CreateBookShortage(bookShortage *models.BookShortage) error {
	err := utils.DB.Create(bookShortage).Error
	if err != nil {
		return err
	}
	return nil
}
