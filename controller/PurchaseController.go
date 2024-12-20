package controller

import (
	"InterLibrarySystem/models"
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var purchaseService = service.PurchaseService{
	Repo: &repository.PurchaseRepository{},
}

var supplierServiceUsedByPurchase = service.SupplierService{
	Repo: &repository.SupplierRepository{},
}

// CreatePurchase 创建购书单
func CreatePurchase(c *gin.Context) {
	var purchase models.Purchase
	err := c.BindJSON(&purchase)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	//获取外键supplierID
	supplier, err := supplierServiceUsedByPurchase.GetSupplierIDByPhoneNum(purchase.PhoneNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	purchase.SupplierID = supplier.ID
	err = purchaseService.CreatePurchase(&purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "创建成功",
	})

}
