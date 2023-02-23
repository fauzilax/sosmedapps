package services

import (
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"sosmedapps/features/user"
	"sosmedapps/helper"
	"sosmedapps/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("success creating account", func(t *testing.T) {
		inputData := user.Core{Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: "alif342"}
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123"}
		repo.On("Register", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.UserName, res.UserName)
		repo.AssertExpectations(t)
	})

	t.Run("all input must fill", func(t *testing.T) {
		inputData := user.Core{
			Name:     "alif",
			Email:    "alif@example.com",
			Password: "alif342",
		}
		srv := New(repo)
		inputData.Password = "alif342"
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not allowed empty")
		assert.Equal(t, uint(0), res.ID)
	})

	t.Run("internal server error", func(t *testing.T) {
		inputData := user.Core{Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: "alif342"}
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123"}
		repo.On("Register", mock.Anything).Return(resData, errors.New("internal server error")).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("login success", func(t *testing.T) {
		inputEmail := "alif@example.com"
		hashed, _ := helper.GeneratePassword("alif342")
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "alif342")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		inputEmail := "alif@example.com"

		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("not found")).Once()

		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "alif342")
		assert.NotNil(t, token)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	// wrong password
	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "alif@example.com"
		hashed, _ := helper.GeneratePassword("EWQTEQF")
		resData := user.Core{ID: uint(1), Email: "alif@example.com", UserName: "alif123", Password: hashed}
		repo.On("Login", inputEmail).Return(resData, nil).Once()

		srv := New(repo)
		_, res, err := srv.Login(inputEmail, "alif342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "password")
		assert.Empty(t, nil)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123"}

	t.Run("success show profile", func(t *testing.T) {

		repo.On("Profile", 1).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData, res)
		repo.AssertExpectations(t)
	})

	t.Run("account not found", func(t *testing.T) {
		repo.On("Profile", int(4)).Return(user.Core{}, errors.New("data not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(4)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})

	// internal server error
	t.Run("internal server error", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("internal server error")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)
	filePath := filepath.Join("..", "..", "..", "ERD.png")
	// imageFalse, _ := os.Open(filePath)
	// imageFalseCnv := &multipart.FileHeader{
	// 	Filename: imageFalse.Name(),
	// }
	imageTrue, err := os.Open(filePath)
	if err != nil {
		log.Println(err.Error())
	}
	imageTrueCnv := &multipart.FileHeader{
		Filename: imageTrue.Name(),
	}
	inputData := user.Core{ID: 1, Name: "Alif", UserName: "alif123", Image: "ERD.png"}
	resData := user.Core{ID: 1, Name: "Alif", UserName: "alif123", Image: imageTrueCnv.Filename}
	t.Run("success updating account", func(t *testing.T) {
		repo.On("Update", 1, inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(*imageTrueCnv, pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("fail updating account", func(t *testing.T) {
		repo.On("Update", 1, inputData).Return(user.Core{}, errors.New("query error,update fail")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Update(*imageTrueCnv, pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("deleting account successful", func(t *testing.T) {
		repo.On("Delete", 1).Return(nil).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	// internal server error, account fail to delete
	t.Run("internal server error, account fail to delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(errors.New("internal server error, account fail to delete")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "fail")
		repo.AssertExpectations(t)
	})

	t.Run("delete account fail", func(t *testing.T) {
		repo.On("Delete", 2).Return(errors.New("delete account fail")).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(2)
		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "fail")
		repo.AssertExpectations(t)
	})

	t.Run("jwt not found", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(0)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		err := srv.Delete(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		repo.AssertExpectations(t)
	})
}

func TestLogout(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("success logout", func(t *testing.T) {
		srv := New(repo)
		_, err := srv.Logout()
		assert.Nil(t, err)
		// assert.ErrorContains(t, err, "sucess logout")
		// repo.AssertExpectations(t)
	})
}
func TestSearching(t *testing.T) {
	repo := mocks.NewUserData(t)
	resData := []user.Core{{ID: 1, Name: "eko", UserName: "koe"}}
	t.Run("success Found", func(t *testing.T) {
		repo.On("Searching", "eko").Return(resData, nil)
		srv := New(repo)
		res, err := srv.Searching("eko")
		assert.Nil(t, err)
		assert.Equal(t, resData[0].Name, res[0].Name)
		repo.AssertExpectations(t)
	})
	t.Run("Not found", func(t *testing.T) {
		repo.On("Searching", "").Return([]user.Core{}, errors.New("no user found"))
		srv := New(repo)
		res, err := srv.Searching("")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, []user.Core{}, res)
		repo.AssertExpectations(t)
	})
}

// func TestSearching(t *testing.T) {
// 	repo := mocks.NewUserData(t)
// 	t.Run("account not found", func(t *testing.T) {
// 		inputUserName := "alif123"
// 		resData := user.Core{ID: uint(1), Name: "Alif", UserName: "alif123"}
// 		repo.On("Searching", inputUserName).Return(user.Core{}, errors.New("not found")).Once()

// 		srv := New(repo)
// 		token, res, err := srv.Searching(inputUserName, resData)
// 		assert.NotNil(t, token)
// 		assert.ErrorContains(t, err, "not found")
// 		assert.Empty(t, token)
// 		assert.Equal(t, uint(0), res.ID)
// 		repo.AssertExpectations(t)
// 	})
// }
