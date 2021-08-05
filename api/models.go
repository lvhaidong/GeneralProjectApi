package api

import "time"

type BaseModel struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedAt *time.Time `gorm:"comment:'创建时间'" json:"crate_at,omitempty"`
	UpdatedAt *time.Time `gorm:"comment:'更新时间'" json:"updated_at,omitempty"`
	DeletedAt *time.Time `gorm:"comment:'删除时间'" sql:"index" json:"delete_at,omitempty"`
	DeletedOn  int `gorm:"default:0;comment:'删除标识'" sql:"index" json:"deleted_on"`
}
