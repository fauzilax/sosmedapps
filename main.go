package main

import (
	"log"
	"sosmedapps/config"
	comData "sosmedapps/features/comment/data"
	comHdl "sosmedapps/features/comment/handler"
	comSrv "sosmedapps/features/comment/services"
	contData "sosmedapps/features/contents/data"
	contHdl "sosmedapps/features/contents/handler"
	contSrv "sosmedapps/features/contents/services"
	uData "sosmedapps/features/user/data"
	uHdl "sosmedapps/features/user/handler"
	uSrv "sosmedapps/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	// gunakan migrate
	config.Migrate(db)

	// gunakan New dari method user
	usrData := uData.New(db)
	usrSrv := uSrv.New(usrData)
	usrHdl := uHdl.New(usrSrv)

	cData := contData.New(db)
	cSrv := contSrv.New(cData)
	cHdl := contHdl.New(cSrv)

	cmData := comData.New(db)
	cmSrv := comSrv.New(cmData)
	cmHdl := comHdl.New(cmSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	// USER
	e.POST("/register", usrHdl.Register())
	e.POST("/login", usrHdl.Login())
	e.PUT("/users", usrHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users", usrHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/users", usrHdl.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/search", usrHdl.Searching())

	e.GET("/logout", usrHdl.Logout(), middleware.JWT([]byte(config.JWTKey)))

	//CONTENT
	e.POST("/contents", cHdl.AddContent(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/contents", cHdl.AllContent())
	e.GET("/contents/:id", cHdl.DetailContent())
	e.PUT("/contents/:id", cHdl.UpdateContent(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/contents/:id", cHdl.DeleteContent(), middleware.JWT([]byte(config.JWTKey)))

	// COMMENT
	e.POST("/comments/:id", cmHdl.NewComment(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/comments/:id", cmHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))
	// e.GET("/comments/all", cmHdl.GetCom())

	// e.POST("/remote", helper.RemoteUpload)
	// ========== Run Program ===========
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
