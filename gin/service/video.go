package service

import "github.com/Akshit8/gin/entity"

// VideoService interface
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

// New ...
func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
