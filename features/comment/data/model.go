package data

import (
	"sosmedapps/features/comment"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comment   string
	UserID    uint
	ContentID uint
	User      User    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Content   Content `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

type User struct {
	gorm.Model
	Name     string
	UserName string
	Image    string
	Comment  []Comment
}

type Content struct {
	gorm.Model
}

func DataToCore(data Comment) comment.Core {
	return comment.Core{
		ID:        data.ID,
		Comment:   data.Comment,
		ContentID: data.ContentID,
	}
}

func CoreToData(core comment.Core) Comment {
	return Comment{
		Model:     gorm.Model{ID: core.ID},
		Comment:   core.Comment,
		ContentID: core.ContentID,
	}
}
