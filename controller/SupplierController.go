package controller

import (
	"InterLibrarySystem/repository"
	"InterLibrarySystem/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var supplierService = service.SupplierService{
	Repo: &repository.SupplierRepository{},
}

// ShowSupplierInfo 展示供应商信息
func ShowSupplierInfo(c *gin.Context) {
	supplier, err := supplierService.ShowSupplierInfo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     1,
		"supplier": supplier,
	})
}
