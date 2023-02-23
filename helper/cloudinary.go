package helper

import (
	"context"
	"mime/multipart"
	"net/http"
	config "sosmedapps/config"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//create cloudinary instance
	cld, err := cloudinary.NewFromParams(config.CLOUDINARY_CLOUD_NAME, config.CLOUDINARY_API_KEY, config.CLOUDINARY_API_SECRET)
	if err != nil {
		return "", err
	}

	//upload file
	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: config.CLOUDINARY_UPLOAD_FOLDER})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

// --------------Helper Model ----------------
type File struct {
	File multipart.File `json:"file,omitempty" validate:"required"`
}

type Url struct {
	Url string `json:"url,omitempty" validate:"required"`
}

// ----------------Helper Service ------------------
var (
	validate = validator.New()
)

type mediaUpload interface {
	FileUpload(file File) (string, error)
	RemoteUpload(url Url) (string, error)
}
type media struct{}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func (*media) FileUpload(file File) (string, error) {
	//validate
	err := validate.Struct(file)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, err := ImageUploadHelper(file.File)
	if err != nil {
		return "", err
	}
	return uploadUrl, nil
}

func (*media) RemoteUpload(url Url) (string, error) {
	//validate
	err := validate.Struct(url)
	if err != nil {
		return "", err
	}

	//upload
	uploadUrl, errUrl := ImageUploadHelper(url.Url)
	if errUrl != nil {
		return "", err
	}
	return uploadUrl, nil
}

// -----------DTOS Helper---------------
type MediaDto struct {
	StatusCode int       `json:"statusCode"`
	Message    string    `json:"message"`
	Data       *echo.Map `json:"data"`
}

// --------------Helper Controller---------------
// func FileUpload(c echo.Context) error {
// 	//upload
// 	formHeader, err := c.FormFile("file")
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 			MediaDto{
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "error",
// 				Data:       &echo.Map{"data": "Select a file to upload"},
// 			})
// 	}

// 	//get file from header
// 	formFile, err := formHeader.Open()
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 			MediaDto{
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "error",
// 				Data:       &echo.Map{"data": err.Error()},
// 			})
// 	}

// 	uploadUrl, err := NewMediaUpload().FileUpload(File{File: formFile})
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError,
// 			MediaDto{
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "error",
// 				Data:       &echo.Map{"data": err.Error()},
// 			})
// 	}

// 	return c.JSON(http.StatusOK,
// 		MediaDto{
// 			StatusCode: http.StatusOK,
// 			Message:    "success",
// 			Data:       &echo.Map{"data": uploadUrl},
// 		})
// }

func RemoteUpload(c echo.Context) error {
	var url Url

	//validate the request body
	if err := c.Bind(&url); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			MediaDto{
				StatusCode: http.StatusBadRequest,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
	}

	uploadUrl, err := NewMediaUpload().RemoteUpload(url)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Error uploading file"},
			})
	}

	return c.JSON(http.StatusOK,
		MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}
