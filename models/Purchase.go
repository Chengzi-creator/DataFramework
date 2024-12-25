package models

type Purchase struct {
	ID          int      `gorm:"primaryKey" json:"id"`                                                         // 自动生成的ID
	Name        string   `gorm:"type:varchar(30);not null" json:"name"`                                        // 供应商名称
	PhoneNumber string   `gorm:"type:varchar(11);not null" json:"phone_number"`                                // 电话号码
	SupplyInfo  string   `gorm:"type:text" json:"supply_info"`                                                 // 供货信息
	SupplierID  int      `gorm:"not null" json:"supplier_id"`                                                  // 外键，关联到Supplier模型
	Supplier    Supplier `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"` // 外键关联
}

func (Purchase) TableName() string {
	return "purchase"
}
