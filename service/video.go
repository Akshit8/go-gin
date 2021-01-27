// Package service implements videos reource service
package service

import "github.com/Akshit8/go-gin/entity"

// VideoService interface defines availaible methods
type VideoService interface {
	Save(*entity.Video) *entity.Video
	FindAll() []*entity.Video
}

type videoService struct {
	videos []*entity.Video
}

// NewVideoService inits videService
func NewVideoService() VideoService {
	return &videoService{
		videos: []*entity.Video{},
	}
}

func (svc *videoService) Save(video *entity.Video) *entity.Video {
	svc.videos = append(svc.videos, video)
	return video
}

func (svc *videoService) FindAll() []*entity.Video {
	return svc.videos
}