package main

import (
	"github.com/Akshit8/gin/controller"
	"github.com/Akshit8/gin/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func main() {
	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	// enabling cors
	server.Use(cors.Default())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(":8000")
}
