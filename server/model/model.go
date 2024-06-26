package model

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	// Id 主键Id
	Id uint `gorm:"primarykey" json:"id"`
	// CreatedAt 创建时间
	CreatedAt time.Time `json:"created_at"`
	// UpdatedAt 更新时间
	UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt 删除时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
