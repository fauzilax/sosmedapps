package handler

import (
	"sosmedapps/features/user"

	"gorm.io/gorm"
)

type Register struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"username"`
}

func RegisterResponse(data user.Core) Register {
	return Register{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		UserName: data.UserName,
	}
}

type Login struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func LoginResponse(data user.Core) Login {
	return Login{
		ID:       data.ID,
		Name:     data.Name,
		UserName: data.UserName,
	}
}

type Search struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func SearchResponse(data user.Core) Search {
	return Search{
		ID:       data.ID,
		Name:     data.Name,
		UserName: data.UserName,
	}
}

type User struct {
	gorm.Model
	Name     string `json:"name" from:"name"`
	Email    string `json:"email" from:"email"`
	Bio      string `json:"bio" from:"bio"`
	Image    string `json:"profilepicture" from:"profilepicture"`
	UserName string `json:"username" from:"username"`
	Password string `json:"password" from:"password"`
	Content  ContentCore
}
type ContentCore struct {
	ID           uint   `json:"id_content" from:"id_content"`
	Content      string `json:"content" from:"content"`
	ContentImage string `json:"image" from:"image"`
	CreateAt     string `json:"create_at" from:"create_at"`
	NumbComment  uint   `json:"comments" from:"comments"`
}

func ProfileResponse(data User) interface{} {
	return User{
		Model:    gorm.Model{ID: data.ID},
		UserName: data.UserName,
		Name:     data.Name,
		Bio:      data.Bio,
		Image:    data.Image,
		Content: ContentCore{
			ID:           data.Content.ID,
			Content:      data.Content.Content,
			ContentImage: data.Content.ContentImage,
			CreateAt:     data.Content.CreateAt,
			NumbComment:  data.Content.NumbComment,
		},
	}
}
