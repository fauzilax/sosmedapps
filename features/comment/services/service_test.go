package services

import (
	"errors"
	"sosmedapps/features/comment"
	"sosmedapps/helper"
	"sosmedapps/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestNewComment(t *testing.T) {
	data := mocks.NewCommentData(t)

	resData := comment.Core{ID: 1, Comment: "yoi"}
	t.Run("success creating comment", func(t *testing.T) {
		data.On("NewComment", 1, uint(1), "yoi").Return(resData, nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.NewComment(id, 1, "yoi")
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.Comment, res.Comment)
		data.AssertExpectations(t)
	})
	t.Run("fail creating comment", func(t *testing.T) {
		data.On("NewComment", 1, uint(1), "yoi").Return(comment.Core{}, errors.New("server error")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		res, err := srv.NewComment(id, 1, "yoi")
		assert.NotNil(t, err)
		assert.Equal(t, "", res.Comment)
		data.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	data := mocks.NewCommentData(t)
	t.Run("success delete", func(t *testing.T) {
		data.On("Delete", uint(1), uint(1)).Return(nil).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		err := srv.Delete(id, 1)
		assert.Nil(t, err)
		data.AssertExpectations(t)
	})
	t.Run("fail deleting", func(t *testing.T) {
		data.On("Delete", uint(1), uint(1)).Return(errors.New("server error")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		err := srv.Delete(id, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "error")
		data.AssertExpectations(t)
	})
	t.Run("wrong deleting", func(t *testing.T) {
		data.On("Delete", uint(1), uint(1)).Return(errors.New("comment cannot deleted")).Once()
		srv := New(data)
		_, tokenIDUser := helper.GenerateToken(1)
		id := tokenIDUser.(*jwt.Token)
		id.Valid = true
		err := srv.Delete(id, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not allowed")
		data.AssertExpectations(t)
	})
}
