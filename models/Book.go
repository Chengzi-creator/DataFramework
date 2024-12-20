package models

type Book struct {
	ID            int            `gorm:"primaryKey" json:"id"`                                                    // 自动生成的ID
	Name          string         `gorm:"type:varchar(150);not null;index:idx_name_publish,unique" json:"name"`    // 书名,复合唯一约束：确保每本书唯一
	Publish       string         `gorm:"type:varchar(150);not null;index:idx_name_publish,unique" json:"publish"` // 出版社名称,复合唯一约束：确保每本书唯一
	Time          string         `gorm:"type:varchar(150);not null" json:"time"`                                  // 出版时间（整数类型）
	Price         float64        `gorm:"type:decimal(6,2);not null" json:"price"`                                 // 价格，最多6位数字，2位小数
	Keyword       string         `gorm:"type:varchar(150)" json:"keyword"`                                        // 关键词
	Stock         int            `gorm:"default:0" json:"stock"`                                                  // 库存量，默认为0
	Supplier      string         `gorm:"type:varchar(10)" json:"supplier"`                                        // 供应商信息
	SeriesNo      int            `gorm:"not null" json:"series_no"`                                               // 系列编号
	Writer        string         `gorm:"writer" json:"writer"`                                                    // 作者
	StoreLocation string         `gorm:"type:varchar(150);not null" json:"store_location"`                        // 存储位置
	BookShortages []BookShortage `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;" json:"-"`                 // 书籍短缺的外键关系
}

func (Book) TableName() string {
	return "book"
}
