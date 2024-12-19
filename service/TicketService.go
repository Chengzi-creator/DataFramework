package service

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"errors"
)

type TicketService struct {
	Repo *repository.TicketRepository
}

// GetTicketsByUserID 根据用户ID获取订单
func (s *TicketService) GetTicketsByUserID(userID int) ([]models.Ticket, error) {
	tickets, err := s.Repo.FindTicketsByUserID(userID)
	if err != nil {
		return nil, errors.New("订单查询失败")
	}
	return tickets, nil
}

// CreateTicket 创建订单
func (s *TicketService) CreateTicket(ticket models.Ticket) error {
	// 调用 Repository 保存订单
	err := s.Repo.CreateTicket(ticket)
	if err != nil {
		return errors.New("订单创建失败")
	}
	return nil
}
