package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Akshit8/go-gin/controller"
	"github.com/Akshit8/go-gin/middleware"
	"github.com/Akshit8/go-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)
)

func setLogOutput() {
	f, _ := os.Create("log/app.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()

	setLogOutput()

	// recovers from panic and return 500
	server.Use(gin.Recovery())
	// dumps http header/body for both request and response
	// server.Use(gindump.Dump())

	// need a new impl of gin server to over write existing
	// logger format
	server.Use(middleware.Logger())

	// add on middleware
	// can attach on default gin server
	// provides basic auth func
	// server.Use(middleware.BasicAuth())

	server.Static("/css", "./template/css")
	server.LoadHTMLGlob("templates/*.html")

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "saved video successfully",
				})
			}
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	server.Run(port)
}
