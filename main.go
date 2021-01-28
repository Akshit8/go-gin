package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Akshit8/go-gin/controller"
	"github.com/Akshit8/go-gin/middleware"
	"github.com/Akshit8/go-gin/repository"
	"github.com/Akshit8/go-gin/service"
	"github.com/gin-gonic/gin"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository("sqlite/main.db")
	videoService    service.VideoService       = service.NewVideoService(videoRepository)
	jwtService      service.JWTService         = service.NewJWTService()
	loginService    service.LoginService       = service.NewLoginService()

	videoController controller.VideoController = controller.NewVideoController(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
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

	server.Static("/css", "./ui/css")
	server.LoadHTMLGlob("ui/*.html")

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "oops! something went wrong",
			})
		}
	})

	// The "/view" endpoints are public (no Authorization required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middleware.AuthorizeJWT())
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

		apiRoutes.GET("/videos/:id", func(ctx *gin.Context) {
			video, err := videoController.Get(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H {
					"video": video,
				})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H {
					"video": "video updated successfully",
				})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadGateway, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusOK, gin.H {
					"video": "video deleted successfully",
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
