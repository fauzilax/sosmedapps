package services

import (
	"errors"
	"log"
	"sosmedapps/features/comment"
	"sosmedapps/helper"
	"strings"
)

type commentServiceCase struct {
	qry comment.CommentData
}

func New(cd comment.CommentData) comment.CommentService {
	return &commentServiceCase{
		qry: cd,
	}
}

// NewComment implements comment.CommentService
func (css *commentServiceCase) NewComment(token interface{}, contentID uint, NewComment string) (comment.Core, error) {
	id := helper.ExtractToken(token)
	res, err := css.qry.NewComment(id, contentID, NewComment)
	if err != nil {
		log.Println("query error", err.Error())
		return comment.Core{}, errors.New("server error, cannot query data")
	}
	return res, nil
}

// Delete implements comment.CommentService
func (css *commentServiceCase) Delete(token interface{}, commentID uint) error {
	userID := helper.ExtractToken(token)
	err := css.qry.Delete(uint(userID), commentID)
	if err != nil {
		log.Println("query error", err.Error())
		if strings.Contains(err.Error(), "cannot") {
			return errors.New("you are not allowed delete other people comment")
		}
		return errors.New("server error")
	}
	return nil
}

// // GetCom implements comment.CommentService
// func (csc *commentServiceCase) GetCom() ([]comment.Core, error) {
// 	res, err := csc.qry.GetCom()
// 	if err != nil {
// 		log.Println("query error", err.Error())
// 		return []comment.Core{}, errors.New("server error, cannot query data")
// 	}
// 	return res, nil
// }
