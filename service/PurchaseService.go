package service

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
)

type PurchaseService struct {
	Repo *repository.PurchaseRepository
}

// CreatePurchase 创建购书单
func (s *PurchaseService) CreatePurchase(purchase *models.Purchase) error {
	err := s.Repo.CreatePurchase(purchase)
	if err != nil {
		return err
	}
	return nil
}
