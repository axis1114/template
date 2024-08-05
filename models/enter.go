package models

import "time"

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type MODEL struct {
	ID        uint      `gorm:"primaryKey;comment:id" json:"id" structs:"-"`
	CreatedAt time.Time `gorm:"comment:创建时间" json:"created_at" structs:"-"`
	UpdatedAt time.Time `gorm:"comment:更新时间" json:"updated_at" structs:"-"`
}

type RemoveRequest struct {
	IDList []uint `json:"id_list"`
}
