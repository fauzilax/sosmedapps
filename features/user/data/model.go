package data

import (
	"sosmedapps/features/contents/data"
	"sosmedapps/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Bio      string
	Image    string
	UserName string
	Password string
	Content  []data.Content `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// type Content struct {
// 	gorm.Model
// 	Content      string
// 	ContentImage string
// 	CreateAt     string
// 	NumbComment  uint
// 	User         User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
// }

func DataToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Bio:      data.Bio,
		Image:    data.Image,
		UserName: data.UserName,
		Password: data.Password,
		Content:  []user.ContentCore{},
	}
}

func CoreToData(core user.Core) User {
	return User{
		Model:    gorm.Model{ID: core.ID},
		Name:     core.Name,
		Email:    core.Email,
		Bio:      core.Bio,
		Image:    core.Image,
		UserName: core.UserName,
		Password: core.Password,
		Content:  []data.Content{},
	}
}
