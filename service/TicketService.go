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
