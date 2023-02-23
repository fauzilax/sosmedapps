package handler

import (
	"log"
	"net/http"
	"sosmedapps/features/comment"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type commentController struct {
	srv comment.CommentService
}

func New(ch comment.CommentService) comment.CommentHandler {
	return &commentController{
		srv: ch,
	}
}

// NewComment implements comment.CommentHandler
func (cc *commentController) NewComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		cID := c.Param("id")
		contentID, _ := strconv.Atoi(cID)
		input := NewComment{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "comment cannot allowed empty"})
		}
		res, err := cc.srv.NewComment(c.Get("user"), uint(contentID), input.Comment)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		log.Println(res)
		return c.JSON(http.StatusCreated, map[string]interface{}{
			// "data":    res,
			"message": "success create comment",
		})
	}
}

// Delete implements comment.CommentHandler
func (cc *commentController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		cID := c.Param("id")
		commentID, _ := strconv.Atoi(cID)
		err := cc.srv.Delete(c.Get("user"), uint(commentID))
		if err != nil {
			if strings.Contains(err.Error(), "not") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you are not allowed delete other people comment"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error, deleting comment fail"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "delete comment success",
		})
	}
}

// // GetCom implements comment.CommentHandler
// func (cc *commentController) GetCom() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		res, err := cc.srv.GetCom()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
// 		}
// 		return c.JSON(http.StatusCreated, map[string]interface{}{
// 			"data":    res,
// 			"message": "success view all comment",
// 		})
// 	}
// }
