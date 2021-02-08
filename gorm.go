package util

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const (
	maxSize     = 1000
	defaultSize = 20
	defaultPage = 1
)

type PaginationReq struct {
	Search   string `form:"search" json:"search"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"page_size" json:"page_size"`
}

type PaginationRes struct {
	PageSize int         `form:"page_size" json:"page_size"`
	Page     int         `form:"page" json:"page"`
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
}

func (p *PaginationReq) check() {
	if p.PageSize <= 0 {
		if p.PageSize == -1 {
			p.PageSize = maxSize
		} else {
			p.PageSize = defaultSize
		}
	}
	if p.Page <= 0 {
		p.Page = defaultPage
	}
}

func (p *PaginationReq) Offset() int {
	p.check()
	return (p.Page - 1) * p.PageSize
}

func (p *PaginationReq) Limit() int {
	p.check()
	return p.PageSize
}
