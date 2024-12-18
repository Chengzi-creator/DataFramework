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
func (s *BookService) GetBookByID(bookID string) (models.Book, error) {
	book, err := s.Repo.FindBookByID(bookID)
	if err != nil {
		return models.Book{}, errors.New("书籍不存在或查询失败")
	}
	return book, nil
}
