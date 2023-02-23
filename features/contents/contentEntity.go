package contents

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CoreContent struct {
	ID           uint   `json:"id" from:"id"`
	Content      string `validate:"required" json:"content" from:"content"`
	ContentImage string `json:"content_image" from:"content_image"`
	CreateAt     string `json:"create_at" from:"create_at"`
	NumbComment  uint   `json:"comment" from:"comment"`
	Users        CoreUser
	Comment      []CommentCore
}

type CoreUser struct {
	ID       uint   `json:"id" from:"id"`
	UserName string `json:"username" from:"username"`
	Name     string `json:"name" from:"name"`
	Image    string `json:"profilepicture" from:"profilepicture"`
}

type CommentCore struct {
	ID       uint
	UserName string
	Comment  string
	Content  CoreContent
}

type ContentHandler interface {
	AddContent() echo.HandlerFunc
	UpdateContent() echo.HandlerFunc
	DetailContent() echo.HandlerFunc
	AllContent() echo.HandlerFunc
	DeleteContent() echo.HandlerFunc
}

type ContentService interface {
	AddContent(token interface{}, fileHeader multipart.FileHeader, newContent CoreContent) (CoreContent, error)
	UpdateContent(token interface{}, contentID uint, content string) (string, error)
	DetailContent(contentID uint) (interface{}, error)
	AllContent() ([]CoreContent, error)
	DeleteContent(token interface{}, contentID uint) error
}

type ContentData interface {
	AddContent(userID uint, newContent CoreContent) (CoreContent, error)
	UpdateContent(userID uint, contentID uint, content string) (string, error)
	DetailContent(contentID uint) (interface{}, error)
	AllContent() ([]CoreContent, error)
	DeleteContent(userID uint, contentID uint) error
}
