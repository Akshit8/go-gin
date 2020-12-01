package controller

import (
	"github.com/Akshit8/gin/entity"
	"github.com/Akshit8/gin/service"
	"github.com/Akshit8/gin/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// VideoController interface
type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

// New ...
func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterAlias("titleCool", validators.ValidateCoolTitle)
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}