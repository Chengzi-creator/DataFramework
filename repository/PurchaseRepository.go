package repository

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
)

type PurchaseRepository struct{}

// CreatePurchase 创建购书单
func (r *PurchaseRepository) CreatePurchase(purchase *models.Purchase) error {
	err := utils.DB.Create(&purchase).Error
	if err != nil {
		return err
	}
	return nil
}
