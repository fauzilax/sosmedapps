package data

import (
	"errors"
	"fmt"
	"log"
	"sosmedapps/features/comment"

	"gorm.io/gorm"
)

type commentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &commentQuery{
		db: db,
	}
}

// NewComment implements comment.CommentData
func (cq *commentQuery) NewComment(id int, contentID uint, newComment string) (comment.Core, error) {
	data := Comment{}
	data.Comment = newComment
	data.ContentID = contentID
	data.UserID = uint(id)
	err := cq.db.Create(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return comment.Core{}, errors.New("server error")
	}
	usrQry := User{}
	err = cq.db.Where("id=?", id).First(&usrQry).Error
	if err != nil {
		return comment.Core{}, errors.New("server error")
	}
	res := comment.Core{}
	res.ID = data.ID
	res.ContentID = contentID
	res.User.ID = usrQry.ID
	res.User.UserName = usrQry.UserName
	res.User.Name = usrQry.Name
	res.Comment = newComment
	res.CreateAt = fmt.Sprintf("%d - %s - %d", data.CreatedAt.Day(), data.CreatedAt.Month(), data.CreatedAt.Year())
	return res, nil

}

// Delete implements comment.CommentData
func (cq *commentQuery) Delete(userID uint, commentID uint) error {
	// log.Println(userID, commentID)
	vld := Comment{}
	err := cq.db.Where("user_id=? AND id=?", userID, commentID).First(&vld).Error
	if err != nil {
		log.Println("comment not found", err.Error())
		return errors.New("comment cannot deleted")
	}
	qry := cq.db.Delete(&Comment{}, commentID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no comment has delete")
	}
	err = qry.Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("delete comment fail")
	}
	return nil
}

// // GetCom implements comment.CommentData
// func (cq commentQuery) GetCom() ([]comment.Core, error) {
// 	qry := []Comment{}
// 	res := cq.db.Preload("User").Find(&qry)
// 	if res.Error != nil {
// 		return []comment.Core{}, errors.New("server error")
// 	}
// 	hasil := []comment.Core{}
// 	for i := 0; i < len(qry); i++ {
// 		hasil = append(hasil, DataToCore(qry[i]))
// 		usrQry := User{}
// 		err := cq.db.Where("id=?", qry[0].UserID).First(&usrQry).Error
// 		if err != nil {
// 			return []comment.Core{}, errors.New("server error")
// 		}
// 		hasil[i].User.ID = usrQry.ID
// 		hasil[i].User.UserName = usrQry.UserName
// 		hasil[i].User.Name = usrQry.Name
// 	}
// 	return hasil, nil
// }
