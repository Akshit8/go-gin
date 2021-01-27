// Package repository implements persistence for api
package repository

import (
	"github.com/Akshit8/go-gin/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// VideoRepository impls video repo interface
type VideoRepository interface {
	Save(video entity.Video)
	Get(video entity.Video)
	FindAll() []entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
}

//TODO: impl CloseDB method acc. to new gorm lib

type database struct {
	connection *gorm.DB
}

// NewVideoRepository inits a new VideoRepository
func NewVideoRepository() VideoRepository {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	db.AutoMigrate(&entity.Video{}, &entity.Person{})
	return &database{
		connection: db,
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Get(video entity.Video) {
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video, video.URL)
}
