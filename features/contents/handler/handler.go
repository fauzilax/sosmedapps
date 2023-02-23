package handler

import (
	"log"
	"net/http"
	"sosmedapps/features/contents"
	"sosmedapps/helper"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type contentController struct {
	srv contents.ContentService
}

func New(cs contents.ContentService) contents.ContentHandler {
	return &contentController{
		srv: cs,
	}
}

// AddContent implements contents.ContentHandler
func (cc *contentController) AddContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}
		input := AddContentRequest{}
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		// Proses Input Ke Service
		res, err := cc.srv.AddContent(c.Get("user"), *formHeader, *RequstToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "type") {
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
			"message": "success create content",
		})
	}
}

// AllContent implements contents.ContentHandler
func (cc *contentController) AllContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cc.srv.AllContent()
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		result := []AllContent{}
		for i := 0; i < len(res); i++ {
			result = append(result, AllContentResponse(res[i]))
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    result,
			"message": "success",
		})

	}
}

// DetailContent implements contents.ContentHandler
func (cc *contentController) DetailContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		cID := c.Param("id")
		contID, _ := strconv.Atoi(cID)
		res, err := cc.srv.DetailContent(uint(contID))
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "content not found"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success show detail",
		})

	}
}

// UpdateContent implements contents.ContentHandler
func (cc *contentController) UpdateContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		cID := c.Param("id")
		contID, _ := strconv.Atoi(cID)
		input := EditContentRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		res, err := cc.srv.UpdateContent(c.Get("user"), uint(contID), input.Content)
		if err != nil {
			if strings.Contains(err.Error(), "not") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you are not allowed edited other people content"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success update data",
		})
	}
}

// DeleteContent implements contents.ContentHandler
func (cc *contentController) DeleteContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		cID := c.Param("id")
		contID, _ := strconv.Atoi(cID)
		err := cc.srv.DeleteContent(c.Get("user"), uint(contID))
		if err != nil {
			if strings.Contains(err.Error(), "not") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "you are not allowed delete other people content"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error, deleting content fail"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete content from user post",
		})
	}
}
