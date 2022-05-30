package model

import "time"
import "gorm.io/gorm"
import "github.com/spf13/cast"

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampField 时间戳
type CommonTimestampField struct {
	CreateAt time.Time `gorm:"column:create_at;index;" json:"create_at,omitempty"`
	UpdateAt time.Time `gorm:"column:update_at;index" json:"update_at,omitempty"`
}

// CommonSoftDeletesField 软删除
type CommonSoftDeletesField struct {
	DeleteAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"delete_at,omitempty"`
}

func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}
