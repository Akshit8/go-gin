package main

import (
	"io"
	"os"

	"github.com/Akshit8/go-gin/api"
	"github.com/Akshit8/go-gin/controller"
	"github.com/Akshit8/go-gin/docs"
	"github.com/Akshit8/go-gin/middleware"
	"github.com/Akshit8/go-gin/repository"
	"github.com/Akshit8/go-gin/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
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

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Akshit - Video API"
	docs.SwaggerInfo.Description = "Rest API in golang following best practices, built with gin, gorm(sqlite), swagger and MVC architecture."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

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
	// The "/view" endpoints are public (no Authorization required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	videoAPI := api.NewVideoAPI(loginController, videoController)

	// Login Endpoint: Authentication + Token creation
	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		auth := apiRoutes.Group("/auth")
		{
			auth.POST("/login", videoAPI.Authenticate)
		}
		videos := apiRoutes.Group("/videos", middleware.AuthorizeJWT())
		{
			videos.GET("/", videoAPI.AllVideos)
			videos.POST("/", videoAPI.CreateVideo)
			videos.GET("/:id", videoAPI.GetVideo)
			videos.PUT("/:id", videoAPI.UpdateVideo)
			videos.DELETE("/:id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	server.Run(port)
}
