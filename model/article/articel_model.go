package article

import (
	"github.com/l6l6ng/bud-hello/model"
	"github.com/l6l6ng/bud-hello/model/tag"
	"github.com/l6l6ng/bud-hello/model/user"
)

type Article struct {
	model.BaseModel
	Title    string `gorm:"not null;index:idx_title_sub" json:"title"`
	SubTitle string `gorm:"not null;index:idx_title_sub" json:"sub_title"`
	Content  string `gorm:"text" json:"content"`

	//用户关联
	UserID string    `gorm:"index" json:"user_id"`
	User   user.User `gorm:"user"`

	//标签关联
	TagID string  `gorm:"index" json:"tag_id"`
	Tag   tag.Tag `gorm:"tag"`

	model.CommonTimestampField
	model.CommonSoftDeletesField
}
