package handler

import "sosmedapps/features/contents"

type CoreContent struct {
	ID           uint   `json:"id_content" from:"id"`
	Content      string `validate:"required" json:"content" from:"content"`
	ContentImage string `json:"content_image" from:"content_image"`
	CreateAt     string `json:"create_at" from:"create_at"`
	NumbComment  uint   `json:"comment" from:"comment"`
	Users        CoreUser
	Comment      []CommentCore
}
type CoreUser struct {
	ID       uint   `json:"id_user" from:"id"`
	UserName string `json:"username" from:"username"`
	Name     string `json:"name" from:"name"`
	Image    string `json:"profilepicture" from:"profilepicture"`
}
type CommentCore struct {
	ID       uint   `json:"id_comment" from:"id"`
	UserName string `json:"username" from:"username"`
	Comment  string `json:"comment" from:"comment"`
	Content  CoreContent
}

type AllContent struct {
	ID           uint   `json:"id_content" from:"id_content"`
	Content      string `validate:"required" json:"content" from:"content"`
	ContentImage string `json:"content_image" from:"content_image"`
	CreateAt     string `json:"create_at" from:"create_at"`
	Users        CoreUser
	NumbComment  uint `json:"comment" from:"comment"`
}

func AllContentResponse(data contents.CoreContent) AllContent {
	return AllContent{
		ID:           data.ID,
		Content:      data.Content,
		ContentImage: data.ContentImage,
		CreateAt:     data.CreateAt,
		NumbComment:  data.NumbComment,
		Users:        CoreUser(data.Users),
	}
}

// type Detail struct {
// 	ID        uint                     `mapstructure:"Name"`
// 	Content   string                   `mapstructure:"content"`
// 	Image     string                   `mapstructure:"image"`
// 	Comments  int                      `mapstructure:"comments"`
// 	Create_at string                   `mapstructure:"create_at"`
// 	Users     interface{}              `mapstructure:"users"`
// 	Comment   []map[string]interface{} `mapstructure:"comment"`
// }
