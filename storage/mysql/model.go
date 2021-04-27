package mysql

import (
	"gorm.io/gorm"
)

type Model struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	CreatedTime int64          `gorm:"default:timestamp" json:"created_time"`
	UpdatedTime int64          `gorm:"default:timestamp on update timestamp" json:"updated_time"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
