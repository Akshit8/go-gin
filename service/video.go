// Package service implements video-api servicec
package service

import (
	"github.com/Akshit8/go-gin/entity"
	"github.com/Akshit8/go-gin/repository"
)

// VideoService interface defines availaible methods
type VideoService interface {
	Save(video entity.Video)
	Get(video entity.Video) entity.Video
	FindAll() []entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
}

type videoService struct {
	repository repository.VideoRepository
}

// NewVideoService inits videService
func NewVideoService(videoRepository repository.VideoRepository) VideoService {
	return &videoService{
		repository: videoRepository,
	}
}

func (svc *videoService) Save(video entity.Video) {
	svc.repository.Save(video)
}

func (svc *videoService) FindAll() []entity.Video {
	return svc.repository.FindAll()
}

func (svc *videoService) Get(video entity.Video) entity.Video {
	return svc.repository.Get(video)
}

func (svc *videoService) Update(video entity.Video) {
	svc.repository.Update(video)
}

func (svc *videoService) Delete(video entity.Video) {
	svc.repository.Delete(video)
}
