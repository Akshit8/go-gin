// Package controller implements api endpoint controllers
package controller

import (
	"github.com/Akshit8/go-gin/entity"
	"github.com/Akshit8/go-gin/service"
	"github.com/gin-gonic/gin"
)

// VideoController defines methods availaible
type VideoController interface {
	Save(ctx *gin.Context) *entity.Video
	FindAll() []*entity.Video
}

type controller struct {
	service service.VideoService
}

// NewVideoController inits new controller for resource video
func NewVideoController(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) Save(ctx *gin.Context) *entity.Video {
	var video *entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}

func (c *controller) FindAll() []*entity.Video {
	return c.service.FindAll()
}
