// Package controller implements api endpoint controllers
package controller

import (
	"log"
	"net/http"

	"github.com/Akshit8/go-gin/entity"
	"github.com/Akshit8/go-gin/service"
	"github.com/Akshit8/go-gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// VideoController defines methods availaible
type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []*entity.Video
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

// NewVideoController inits new controller for resource video
func NewVideoController(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("cool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error{
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		log.Print("error in binding")
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		log.Print("error in validation")
		return err
	}
	c.service.Save(&video)
	return nil
}

func (c *controller) FindAll() []*entity.Video {
	return c.service.FindAll()
}

func (c *controller) ShowAll(ctx *gin.Context) {
	videos := c.service.FindAll()
	data := gin.H{
		"title": "Video Page",
		"videos": videos,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
