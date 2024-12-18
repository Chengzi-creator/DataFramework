package models

type BookShortage struct {
	ID               int    `gorm:"primaryKey" json:"id"`                        // 自动生成的ID
	Name             string `gorm:"type:varchar(150);not null" json:"name"`      // 书名
	Publish          string `gorm:"type:varchar(150);not null" json:"publish"`   // 出版社名称
	Supplier         string `gorm:"type:varchar(10);not null" json:"supplier"`   // 供应商信息
	Quantity         int    `gorm:"not null;check:quantity > 0" json:"quantity"` // 短缺数量，验证大于0
	RegistrationDate int    `gorm:"not null" json:"registration_date"`           // 登记日期
	BookID           int    `gorm:"not null" json:"book_id"`
}

func (BookShortage) TableName() string {
	return "book_shortage"
}
