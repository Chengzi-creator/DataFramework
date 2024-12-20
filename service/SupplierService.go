package service

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
)

type SupplierService struct {
	Repo *repository.SupplierRepository
}

// ShowSupplierInfo 展示供应商信息
func (s *SupplierService) ShowSupplierInfo() ([]models.Supplier, error) {
	supplier, err := s.Repo.ShowSupplier()
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

// GetSupplierIDByPhoneNum 根据供应商电话获取供应商ID，用于purchase外键
func (s *SupplierService) GetSupplierIDByPhoneNum(phoneNumber string) (*models.Supplier, error) {
	supplier, err := s.Repo.GetSupplierIDByPhoneNum(phoneNumber)
	if err != nil {
		return nil, err
	}
	return supplier, nil
}
