package models

type User struct {
	ID           int     `gorm:"primaryKey" json:"user_id"`                               // 自动生成的ID
	Username     string  `gorm:"not null" json:"username"`                                //用户名
	Address      string  `gorm:"type:varchar(300);not null" json:"address"`               // 地址信息
	Balance      float64 `gorm:"type:decimal(10,2);not null" json:"balance"`              // 账户余额，最多10位数字，2位小数
	CreditRating int     `gorm:"not null" json:"credit_rating"`                           // 信用等级
	Password     string  `gorm:"size:255;not null" json:"password"`                       //密码
	IsAdminister bool    `gorm:"type:tinyint(1);not null;default:0" json:"is_administer"` //判断是否为管理员
}

func (User) TableName() string {
	return "user"
}
