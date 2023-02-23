package services

import (
	"errors"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"sosmedapps/features/contents"
	"sosmedapps/helper"
	"sosmedapps/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestAllContent(t *testing.T) {
	data := mocks.NewContentData(t)
	resData := []contents.CoreContent{}
	t.Run("Succes View All Content", func(t *testing.T) {
		data.On("AllContent").Return(resData, nil).Once()
		srv := New(data)
		res, err := srv.AllContent()
		assert.Nil(t, err)
		assert.Equal(t, []contents.CoreContent{}, res)
		data.AssertExpectations(t)
	})
	t.Run("Fail View All Content", func(t *testing.T) {
		data.On("AllContent").Return(resData, errors.New("server error")).Once()
		srv := New(data)
		res, err := srv.AllContent()
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, []contents.CoreContent{}, res)
		data.AssertExpectations(t)
	})
}

func TestAddContent(t *testing.T) {
	data := mocks.NewContentData(t)
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

	input := contents.CoreContent{ID: 1, Content: "ok", ContentImage: imageTrueCnv.Filename}
	resData := contents.CoreContent{ID: 1, Content: "ok", ContentImage: imageTrueCnv.Filename}
	t.Run("Succes Create Content", func(t *testing.T) {
		data.On("AddContent", uint(1), input).Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.AddContent(id, *imageTrueCnv, input)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		data.AssertExpectations(t)

	})
	t.Run("failed Create Content", func(t *testing.T) {
		data.On("AddContent", uint(1), input).Return(contents.CoreContent{}, errors.New("content cannot edited")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.AddContent(id, *imageTrueCnv, input)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		assert.Equal(t, contents.CoreContent{}, res)
		data.AssertExpectations(t)

	})
	// t.Run("fsize error", func(t *testing.T) {
	// 	srv := New(data)
	// 	_, tokenIDUser := helper.GenerateToken(1)
	// 	id := tokenIDUser.(*jwt.Token)
	// 	id.Valid = true
	// 	srv.AddContent(id, *imageTrueCnv, input)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "size error")

	// })
}

func TestDelete(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("success delete content", func(t *testing.T) {
		repo.On("DeleteContent", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteContent(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})
	t.Run("wrong delete", func(t *testing.T) {
		repo.On("DeleteContent", uint(1), uint(1)).Return(errors.New("content cannot deleted")).Once()
		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteContent(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not")
		repo.AssertExpectations(t)
	})
	t.Run("internal server error", func(t *testing.T) {
		repo.On("DeleteContent", uint(1), uint(1)).Return(errors.New("server error")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteContent(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})
}

func TestUpdateContent(t *testing.T) {
	repo := mocks.NewContentData(t)
	// inputData := contents.CoreContent{ID: uint(1), Content: "hello everyone"}
	resData := "hello everybody"

	t.Run("success update content", func(t *testing.T) {
		repo.On("UpdateContent", uint(1), uint(1), "hello everyone").Return(resData, nil).Once()
		srv := New(repo)
		_, tokenIDUser := helper.GenerateToken(1)
		userID := tokenIDUser.(*jwt.Token)
		userID.Valid = true
		_, err := srv.UpdateContent(userID, uint(1), "hello everyone")
		assert.Nil(t, err)
		// assert.Equal(t, resData.ID, userID)
		repo.AssertExpectations(t)
	})
}

//a
