package handler

import (
	"sosmedapps/features/user"
)

type LoginRequest struct {
	UserName string `json:"email" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	UserName string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UpdateRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Image    string `json:"image" form:"image"`
	Bio      string `json:"bio" form:"bio"`
	UserName string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password"`
}

func RequstToCore(dataUser interface{}) *user.Core {
	res := user.Core{}
	switch dataUser.(type) {
	case LoginRequest:
		cnv := dataUser.(LoginRequest)
		res.UserName = cnv.UserName
		res.Password = cnv.Password
	case RegisterRequest:
		cnv := dataUser.(RegisterRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.UserName = cnv.UserName
		res.Password = cnv.Password
	case UpdateRequest:
		cnv := dataUser.(UpdateRequest)
		res.Name = cnv.Name
		res.Email = cnv.Email
		res.Bio = cnv.Bio
		res.Image = cnv.Image
		res.UserName = cnv.UserName
		res.Password = cnv.Password
	default:
		return nil
	}
	return &res

}
