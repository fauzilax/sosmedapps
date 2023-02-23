package data

import (
	cmData "sosmedapps/features/comment/data"
	"sosmedapps/features/contents"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Content      string
	ContentImage string
	NumbComment  uint
	UserID       uint
	User         User
	Comment      []cmData.Comment
}

type User struct {
	gorm.Model
	Name     string
	Image    string
	UserName string
}

func ContentToCore(data Content) contents.CoreContent {
	return contents.CoreContent{
		ID:           data.ID,
		Content:      data.Content,
		ContentImage: data.ContentImage,
		NumbComment:  data.NumbComment,
		Users: contents.CoreUser{
			ID: data.UserID,
		},
		Comment: []contents.CommentCore{},
	}
}

func CoreToData(core contents.CoreContent) Content {
	return Content{
		Model:        gorm.Model{ID: core.ID},
		Content:      core.Content,
		ContentImage: core.ContentImage,
		UserID:       core.Users.ID,
	}
}

// func ToCoreContent(data []Content) []contents.CoreContent {
// 	var tmp []contents.CoreContent
// 	for _, v := range data {
// 		tmp = append(tmp, v.ContentToCore())
// 	}
// 	return tmp
// }
