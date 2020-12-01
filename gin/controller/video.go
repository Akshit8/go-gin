package controller

import (
	"github.com/Akshit8/gin/entity"
	"github.com/Akshit8/gin/service"
	"github.com/gin-gonic/gin"
)

// VideoController interface
type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

// New ...
func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindAll() []entity.Video {
	return c.service.FindAll()
}