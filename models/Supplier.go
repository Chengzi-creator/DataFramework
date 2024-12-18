package models

type Supplier struct {
	ID          int    `gorm:"primaryKey" json:"id"`                         // 自动生成的ID
	Name        string `gorm:"type:varchar(30);not null" json:"name"`        // 供应商名称
	PhoneNumber string `gorm:"type:varchar(11);not null" json:"phoneNumber"` // 电话号码
	SupplyInfo  string `gorm:"type:text" json:"supply_info"`                 // 供货信息
}

func (Supplier) TableName() string {
	return "supplier"
}
