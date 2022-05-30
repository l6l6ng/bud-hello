package user

import "github.com/l6l6ng/bud-hello/model"

type User struct {
	model.BaseModel
	LoginName string `gorm:"type:varchar(255);not null;index;comment:账号" json:"login_name"`
	Name string `gorm:"type:varchar(255);not null;index;comment:用户名"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`

	model.CommonTimestampField
	model.CommonSoftDeletesField
}
