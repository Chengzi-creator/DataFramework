package models

type Store struct {
	ID       int    `gorm:"primaryKey" json:"id"`                                                      // 自动生成的ID
	Location string `gorm:"type:varchar(150);not null" json:"location"`                                // 存储位置
	State    int    `gorm:"default:1" json:"state"`                                                    // 存储状态，默认为1
	BookID   int    `gorm:"not null" json:"book_id"`                                                   // 外键，关联到Book模型
	Book     Book   `gorm:"foreignKey:book_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"` // 外键关联
}

func (Store) TableName() string {
	return "store"
}
