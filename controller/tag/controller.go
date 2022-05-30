package tag

import (
	context "context"
	"gorm.io/gorm"
)

//func Load(db gorm.DB) (*Controller, error) {
//	//return &Controller{
//	//	db
//	//}
//}

type Controller struct {
	DB *gorm.DB
}

// Tag struct
type Tag struct {
	// Fields...
}

// Index of tags
// GET /tag
func (c *Controller) Index(ctx context.Context) (tags []*Tag, err error) {
	return []*Tag{}, nil
}

// Show tag
// GET /tag/:id
func (c *Controller) Show(ctx context.Context, id int) (tag *Tag, err error) {
	return &Tag{}, nil
}
