package handler

import (
	"log"
	"net/http"
	"sosmedapps/features/user"
	"sosmedapps/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	srv user.UserService
}

func New(us user.UserService) user.UserHandler {
	return &userController{
		srv: us,
	}
}

// Register implements user.UserHandler
func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		res, err := uc.srv.Register(*RequstToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "all input must fill"})
			} else if strings.Contains(err.Error(), "username") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username already registered"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			// "data":    RegisterResponse(res),
			"message": "success creating account",
		})

	}
}

// Login implements user.UserHandler
func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		tokenGen, res, err := uc.srv.Login(input.UserName, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "account not registered"})
			} else if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username or password not allowed empty"})
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong password"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    LoginResponse(res),
			"token":   tokenGen,
			"message": "login success",
		})
	}
}

// Profile implements user.UserHandler
func (uc *userController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.srv.Profile(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success show profile",
		})

	}
}

// Delete implements user.UserHandler
func (uc *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uc.srv.Delete(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error, account fail to delete",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "delete data user success",
		})
	}
}

// Update implements user.UserHandler
func (uc *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}
		input := UpdateRequest{}
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		// Proses Input Ke Service
		res, err := uc.srv.Update(*formHeader, c.Get("user"), *RequstToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already used"})
			} else if strings.Contains(err.Error(), "username") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username already used"})
			} else if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "only jpg or png file can be upload"})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "max file size is 500KB"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			// "data":    res,
			"message": "update profile success",
		})
		// if fileType == "image/png" || fileType == "image/jpeg" {
		// } else {
		// 	return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "only jpg or png file can be upload"})
		// }
	}
}

// Searching implements user.UserHandler
func (uc *userController) Searching() echo.HandlerFunc {
	return func(c echo.Context) error {
		quotes := c.QueryParam("q")
		log.Println(quotes)
		res, err := uc.srv.Searching(quotes)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "user not found"})
		}
		result := []Search{}
		for i := 0; i < len(res); i++ {
			result = append(result, SearchResponse(res[i]))
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    result,
			"message": "success mencari data",
		})
	}
}

// Logout implements user.UserHandler
func (uc *userController) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.srv.Logout()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "user not found"})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"token":   res,
			"message": "success logout",
		})
	}
}
