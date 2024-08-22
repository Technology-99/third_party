package database

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at;comment:创建时间;" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"column:updated_at;comment:更新时间;" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type BaseAdminModel struct {
	CreatedBy uint `gorm:"column:created_by;comment:创建者;type: int" json:"createdBy"`
	UpdatedBy uint `gorm:"column:updated_by;comment:更新者;type: int" json:"updatedBy"`
	DeletedBy uint `gorm:"column:deleted_by;comment:删除者;type: int" json:"deletedBy"`
}
