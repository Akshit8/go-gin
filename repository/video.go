// Package repository implements persistence for api
package repository

import (
	"log"

	"github.com/Akshit8/go-gin/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// VideoRepository impls video repo interface
type VideoRepository interface {
	Save(video entity.Video)
	Get(video entity.Video) entity.Video
	FindAll() []entity.Video
	Update(video entity.Video)
	Delete(video entity.Video)
}

//TODO: impl CloseDB method acc. to new gorm lib

type database struct {
	connection *gorm.DB
}

// NewVideoRepository inits a new VideoRepository
func NewVideoRepository(fileName string) VideoRepository {
	db, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = db.AutoMigrate(&entity.Video{}, &entity.Person{})
	if err != nil {
		log.Fatal("failed to migrate db", err)
	}
	return &database{
		connection: db,
	}
}

func (db *database) Save(video entity.Video) {
	log.Print("video obj before repo save: ", video)
	result := db.connection.Create(&video)
	if result.Error != nil {
		log.Print("db save error: ", result.Error.Error())
	}
}

func (db *database) Get(video entity.Video) entity.Video {
	result := db.connection.Preload("Author").First(&video)
	if result.Error != nil {
		log.Print("db get error: ", result.Error.Error())
	}
	return video
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	result := db.connection.Preload("Author").Find(&videos)
	if result.Error != nil {
		log.Print("db find all error: ", result.Error.Error())
	}
	return videos
}

func (db *database) Update(video entity.Video) {
	result := db.connection.Save(&video)
	if result.Error != nil {
		log.Print("db update error: ", result.Error.Error())
	}
}

func (db *database) Delete(video entity.Video) {
	result := db.connection.Delete(&video)
	if result.Error != nil {
		log.Print("db delete error: ", result.Error.Error())
	}
}
