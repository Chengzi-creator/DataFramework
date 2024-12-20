package service

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"errors"
)

// BookService 负责书籍的业务逻辑
type BookService struct {
	Repo *repository.BookRepository
}

// SearchBooks 根据书名查询符合书籍
func (s *BookService) SearchBooks(name string) ([]models.Book, error) {
	// 调用 Repository 层进行查询
	books, err := s.Repo.FindBooksByName(name)
	if err != nil {
		return nil, errors.New("数据库查询失败")
	}

	return books, nil
}

// GetBookByID 根据书籍 ID 获取单个书籍信息
func (s *BookService) GetBookByID(bookID int) (*models.Book, error) {
	book, err := s.Repo.FindBookByID(bookID)
	if err != nil {
		return nil, errors.New("书籍不存在或查询失败")
	}
	return book, nil
}

// GetBookByName 根据书名查询书籍，用于缺书申请查找外键
func (s *BookService) GetBookByName(name string) (*models.Book, error) {
	book, err := s.Repo.FindBookByName(name)
	if err != nil {
		return nil, errors.New("书籍不存在")
	}
	return book, nil
}

// CreateBook 新增书籍
func (s *BookService) CreateBook(book *models.Book) error {
	err := s.Repo.CreateBook(book)
	if err != nil {
		return err
	}
	return nil
}

// UpdateBook 更新书籍信息
func (s *BookService) UpdateBook(book *models.Book) error {
	err := s.Repo.UpdateBook(book)
	if err != nil {
		return err
	}
	return nil
}

// DeleteBookByBookID 删除书籍
func (s *BookService) DeleteBookByBookID(bookID int) error {
	err := s.Repo.DeleteBookByBookID(bookID)
	if err != nil {
		return err
	}
	return nil
}

// ShowBookShortage 显示缺书登记
func (s *BookService) ShowBookShortage() ([]models.BookShortage, error) {
	bookShortage, err := s.Repo.ShowBookShortage()
	if err != nil {
		return nil, err
	}
	return bookShortage, nil
}

// CreateBookShortage 创建缺书记录
func (s *BookService) CreateBookShortage(bookShortage *models.BookShortage) error {
	err := s.Repo.CreateBookShortage(bookShortage)
	if err != nil {
		return err
	}
	return nil
}
