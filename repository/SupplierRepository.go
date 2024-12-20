package repository

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/utils"
)

type SupplierRepository struct{}

// ShowSupplier 显示供应商信息
func (r *SupplierRepository) ShowSupplier() ([]models.Supplier, error) {
	var supplier []models.Supplier
	err := utils.DB.Find(&supplier).Error
	if err != nil {
		return nil, err
	}
	return supplier, nil
}

// GetSupplierIDByPhoneNum  根据供应商电话获取供应商ID，用于purchase外键
func (r *SupplierRepository) GetSupplierIDByPhoneNum(phoneNumber string) (*models.Supplier, error) {
	var supplier models.Supplier
	err := utils.DB.Where("phone_number = ?", phoneNumber).First(&supplier).Error
	if err != nil {
		return nil, err
	}
	return &supplier, nil
}
