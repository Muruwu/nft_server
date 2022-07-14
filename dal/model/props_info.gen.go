// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePropsInfo = "props_info"

// PropsInfo mapped from table <props_info>
type PropsInfo struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"` // 主键id
	PropsID   int64      `gorm:"column:props_id;not null" json:"props_id"`          // 道具id
	Name      string     `gorm:"column:name;not null" json:"name"`                  // 道具名称
	ImageURL  *string    `gorm:"column:image_url" json:"image_url"`                 // 道具图片
	Type      int32      `gorm:"column:type;not null" json:"type"`                  // 道具类型
	Desc      *string    `gorm:"column:desc" json:"desc"`                           // 道具描述
	CreatedAt *time.Time `gorm:"column:created_at" json:"created_at"`               // 创建时间
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`               // 更新时间
}

// TableName PropsInfo's table name
func (*PropsInfo) TableName() string {
	return TableNamePropsInfo
}