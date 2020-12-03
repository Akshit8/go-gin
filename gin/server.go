package main

import (
	"io"
	"net/http"
	"net/url"
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
	// server.Use(middlewares.BasicAuth())
	// shift to api group

	// enabling cors
	server.Use(cors.Default())

	// dump middleware
	// prints extra info to stdout per http req
	server.Use(gindump.Dump())

	// can also club middlewares
	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	// serving static files and html templates
	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.GET("/", func(c *gin.Context) {
		location := url.URL{Path: "/view/videos"}
		c.Redirect(http.StatusFound, location.RequestURI())
	})

	viewRoute := server.Group("/view")
	{
		viewRoute.GET("/videos", videoController.ShowAll)
	}

	// grouping multi-route grouping
	apiRoute := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoute.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoute.POST("/video", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"success": false,
					"error":   err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusCreated, gin.H{
				"success": true,
				"message": "video created successfully",
			})
		})
	}

	port := os.Getenv("PORT")
	server.Run(":" + port)
}
