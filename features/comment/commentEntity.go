package comment

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint   `json:"id" from:"id"`
	Comment   string `json:"comment" from:"comment"`
	ContentID uint   `json:"content_id" from:"content_id"`
	CreateAt  string `json:"create_at" from:"create_at"`
	User      UserCore
}

type UserCore struct {
	ID       uint   `json:"id" from:"id"`
	UserName string `json:"username" from:"username"`
	Name     string `json:"name" from:"name"`
}

type CommentHandler interface {
	NewComment() echo.HandlerFunc
	Delete() echo.HandlerFunc
	// GetCom() echo.HandlerFunc
}

type CommentService interface {
	NewComment(token interface{}, contentID uint, NewComment string) (Core, error)
	Delete(token interface{}, commentID uint) error
	// GetCom() ([]Core, error)
}

type CommentData interface {
	NewComment(userID int, contentID uint, newComment string) (Core, error)
	Delete(userID uint, commentID uint) error
	// GetCom() ([]Core, error)
}
