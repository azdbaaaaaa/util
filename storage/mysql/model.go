package mysql

import (
	"gorm.io/gorm"
)

type Model struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	CreateTime  int64          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  int64          `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
	DeletedTime gorm.DeletedAt `gorm:"index" json:"-"`
}
