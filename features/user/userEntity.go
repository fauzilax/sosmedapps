package user

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

// perjanjian kontrak
type Core struct {
	ID       uint
	Name     string
	Email    string
	Bio      string
	Image    string
	UserName string
	Password string
	Content  []ContentCore
}

type ContentCore struct {
	ID           uint
	Content      string
	ContentImage string
	CreateAt     string
	NumbComment  uint
	User         Core
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Searching() echo.HandlerFunc
	Logout() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
	Profile(userToken interface{}) (interface{}, error)
	Update(formHeader multipart.FileHeader, userToken interface{}, updateData Core) (Core, error)
	Delete(userToken interface{}) error
	Searching(quote string) ([]Core, error)
	Logout() (interface{}, error)
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
	Profile(id int) (interface{}, error)
	Update(id int, updateData Core) (Core, error)
	Delete(id int) error
	Searching(quote string) ([]Core, error)
}
