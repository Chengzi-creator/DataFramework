package repository

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
)

type TicketRepository struct{}

// FindTicketsByUserID 根据用户ID查询订单
func (r *TicketRepository) FindTicketsByUserID(userID int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	err := utils.DB.Where("user_id = ?", userID).Find(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// CreateTicket 创建订单
func (r *TicketRepository) CreateTicket(ticket models.Ticket) error {
	err := utils.DB.Create(&ticket).Error
	if err != nil {
		return err
	}
	return nil
}
