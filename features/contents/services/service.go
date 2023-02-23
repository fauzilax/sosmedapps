package services

import (
	"errors"
	"log"
	"mime/multipart"
	"sosmedapps/features/contents"
	"sosmedapps/helper"
	"strings"
)

type contentServiceCase struct {
	qry contents.ContentData
}

func New(cd contents.ContentData) contents.ContentService {
	return &contentServiceCase{
		qry: cd,
	}
}

// AddContent implements contents.ContentService
func (csc *contentServiceCase) AddContent(token interface{}, formHeader multipart.FileHeader, newContent contents.CoreContent) (contents.CoreContent, error) {
	userID := helper.ExtractToken(token)
	// image prosses
	//validasi size
	if formHeader.Size != 0 {
		if formHeader.Size > 500000 {
			return contents.CoreContent{}, errors.New("size error")
		}
		//get file from header to check type
		formFile, err := formHeader.Open()
		if err != nil {
			return contents.CoreContent{}, errors.New("error open formheader")
		}
		// Validasi Type
		if !helper.TypeFile(formFile) {
			return contents.CoreContent{}, errors.New("file type error")
		}
		defer formFile.Close()
		formFile, _ = formHeader.Open()
		uploadUrl, err := helper.NewMediaUpload().FileUpload(helper.File{File: formFile})
		if err != nil {
			return contents.CoreContent{}, errors.New("server error")
		}
		newContent.ContentImage = uploadUrl
	}
	//input ke query proses
	res, err := csc.qry.AddContent(uint(userID), newContent)
	if err != nil {
		log.Println("query error", err.Error())
		return contents.CoreContent{}, errors.New("server error")
	}
	return res, nil
}

// AllContent implements contents.ContentService
func (csc *contentServiceCase) AllContent() ([]contents.CoreContent, error) {
	res, err := csc.qry.AllContent()
	if err != nil {
		log.Println("query error")
		return []contents.CoreContent{}, errors.New("server error")
	}
	return res, nil
}

// DetailContent implements contents.ContentService
func (csc *contentServiceCase) DetailContent(contentID uint) (interface{}, error) {
	res, err := csc.qry.DetailContent(contentID)
	if err != nil {
		log.Println("query error")
		return contents.CoreContent{}, errors.New("server error")
	}
	return res, nil
}

// UpdateContent implements contents.ContentService
func (csc *contentServiceCase) UpdateContent(token interface{}, contentID uint, content string) (string, error) {
	userID := helper.ExtractToken(token)
	res, err := csc.qry.UpdateContent(uint(userID), contentID, content)
	if err != nil {
		if strings.Contains(err.Error(), "cannot") {
			return "", errors.New("you are not allowed edited other people content")
		}
		return "", errors.New("internal server error")
	}
	return res, nil
}

// DeleteContent implements contents.ContentService
func (csc *contentServiceCase) DeleteContent(token interface{}, contentID uint) error {
	userID := helper.ExtractToken(token)
	err := csc.qry.DeleteContent(uint(userID), contentID)
	if err != nil {
		log.Println("query error")
		if strings.Contains(err.Error(), "cannot") {
			return errors.New("you are not allowed delete other people content")
		}
		return errors.New("server error")
	}
	return nil
}
