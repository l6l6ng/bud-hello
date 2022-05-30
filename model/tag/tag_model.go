package tag

import (
	"github.com/l6l6ng/bud-hello/model"
	"github.com/l6l6ng/bud-hello/model/user"
)

type Tag struct {
	model.BaseModel

	Name string `gorm:"unique" json:"name"`

	//用户关联
	UserId string    `gorm:"index" json:"user_id"`
	User   user.User `gorm:"migration" json:"user"`

	model.CommonTimestampField
	model.CommonSoftDeletesField
}
