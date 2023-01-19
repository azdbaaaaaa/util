package mysql

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID          uint           `gorm:"primary_key" json:"id"`
	CreateTime  time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime  time.Time      `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"update_time"`
	DeletedTime gorm.DeletedAt `gorm:"type:timestamp;index" json:"-"`
}
