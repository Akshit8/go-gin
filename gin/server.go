package main

import (
	"io"
	"os"

	"github.com/Akshit8/gin/controller"
	"github.com/Akshit8/gin/middlewares"
	"github.com/Akshit8/gin/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	// logs to both log files and unix stdout
	// io.MultiWriter - MultiWriter creates a writer that duplicates 
	// its writes to all the provided writers, similar to the Unix tee(1) command.
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

var videoService service.VideoService = service.New()
var videoController controller.VideoController = controller.New(videoService)

func main() {
	// set gin application mode release/debug
	gin.SetMode(gin.DebugMode)
	// setup logger target
	setupLogOutput()

	// for deafult configuration of gin middlewares
	// server := gin.Default()

	// for selef configuration
	server := gin.New()

	//  recover from any panics and writes a 500 if there was one.
	server.Use(gin.Recovery())

	// custom logger middleware
	server.Use(middlewares.Logger())

	// basic middleware auth
	server.Use(middlewares.BasicAuth())

	// enabling cors
	server.Use(cors.Default())

	// dump middleware
	// prints extra info to stdout per http req
	server.Use(gindump.Dump())

	// can also club middlewares
	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/video", func(ctx *gin.Context) {
		ctx.JSON(201, videoController.Save(ctx))
	})

	server.Run(":8000")
}
