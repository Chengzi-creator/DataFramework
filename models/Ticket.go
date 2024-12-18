package models

import "time"

type Ticket struct {
	ID          int       `gorm:"primaryKey" json:"id"`                    // 自动生成的ID
	Price       float64   `gorm:"type:decimal(6,2);not null" json:"price"` // 订单金额，最多6位数字，2位小数
	Time        time.Time `gorm:"type:datetime;not null" json:"time"`      // 创建时间
	Quantity    int       `gorm:"default:1" json:"quantity"`               // 数量，默认为1
	Description string    `gorm:"type:varchar(150)" json:"description"`    // 订单描述，最大150字符
	Address     string    `gorm:"type:varchar(150)" json:"address"`        // 地址信息，最大150字符
	State       int       `gorm:"default:0" json:"state"`                  // 状态，默认为0
	UserID      int       `gorm:"" json:"user_id"`
	User        User      `gorm:"foreignKey:User_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"` // 外键关联
}

func (Ticket) TableName() string {
	return "ticket"
}
