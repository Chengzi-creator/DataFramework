package repository

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
)

type TicketRepository struct{}

// FindTicketsByUserID 根据用户ID查询订单
func (r *TicketRepository) FindTicketsByUserID(userID int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	result := utils.DB.Where("user_id = ?", userID).Find(&tickets)
	return tickets, result.Error
}

// CreateTicket 创建订单
func (r *TicketRepository) CreateTicket(ticket models.Ticket) error {
	return utils.DB.Create(&ticket).Error
}
