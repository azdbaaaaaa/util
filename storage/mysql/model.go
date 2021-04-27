package mysql

import (
	"gorm.io/gorm"
)

type Model struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	CreateTime  int64          `gorm:"default:timestamp" json:"create_time"`
	UpdateTime  int64          `gorm:"default:timestamp on update timestamp" json:"update_time"`
	DeletedTime gorm.DeletedAt `gorm:"index" json:"-"`
}
